package sysTool

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
	modelSysTool "server/internal/model/sysTool"
	serviceSysManagement "server/internal/service/sysManagement"
	"log/slog"
)

type ServiceTokenService struct {
	repo          modelSysTool.ServiceTokenRepository
	casbinService *serviceSysManagement.CasbinService
	ctx           context.Context
}

func NewServiceTokenService() *ServiceTokenService {
	return &ServiceTokenService{
		repo:          modelSysTool.NewServiceTokenRepo(global.TD27_DB),
		casbinService: serviceSysManagement.NewCasbinService(),
		ctx:           context.Background(),
	}
}

func (s *ServiceTokenService) generateToken() (plainToken, tokenHash string) {
	u := uuid.New().String()
	u = strings.ReplaceAll(u, "-", "")
	plainToken = "sk-" + u[:32]

	hash := sha256.Sum256([]byte(plainToken))
	tokenHash = hex.EncodeToString(hash[:])

	return plainToken, tokenHash
}

func (s *ServiceTokenService) Create(req *modelSysTool.CreateServiceTokenReq) (*modelSysTool.CreateServiceTokenResp, error) {
	plainToken, tokenHash := s.generateToken()

	token := &modelSysTool.ServiceToken{
		Name:      req.Name,
		TokenHash: tokenHash,
		Status:    true,
		ExpiresAt: req.ExpiresAt,
	}

	if err := s.repo.Create(s.ctx, token); err != nil {
		return nil, err
	}

	if len(req.ApiIDs) > 0 {
		permissionIDs, err := s.getPermissionIDsByAPIs(req.ApiIDs)
		if err != nil {
			return nil, err
		}

		if err = s.repo.SetTokenPermissions(s.ctx, token.ID, permissionIDs); err != nil {
			return nil, err
		}

		if err = s.syncTokenToCasbin(token.ID, permissionIDs); err != nil {
			slog.Error("同步令牌到Casbin失败", "error", err)
		}
	}

	resp := &modelSysTool.CreateServiceTokenResp{
		Token: plainToken,
	}
	resp.ID = token.ID
	resp.Name = token.Name
	resp.Status = token.Status
	resp.ExpiresAt = token.ExpiresAt
	resp.CreatedAt = token.CreatedAt.Unix()
	resp.ApiCount = len(req.ApiIDs)

	return resp, nil
}

func (s *ServiceTokenService) Update(req *modelSysTool.UpdateServiceTokenReq) error {
	token, err := s.repo.FindByID(s.ctx, req.ID)
	if err != nil {
		return err
	}

	token.Name = req.Name
	token.Status = req.Status
	token.ExpiresAt = req.ExpiresAt

	if err = s.repo.Update(s.ctx, token); err != nil {
		return err
	}

	permissionIDs, err := s.getPermissionIDsByAPIs(req.ApiIDs)
	if err != nil {
		return err
	}

	if err = s.repo.SetTokenPermissions(s.ctx, token.ID, permissionIDs); err != nil {
		return err
	}

	if err = s.syncTokenToCasbin(token.ID, permissionIDs); err != nil {
		slog.Error("同步令牌到Casbin失败", "error", err)
	}

	return nil
}

func (s *ServiceTokenService) Delete(id uint) error {
	subject := fmt.Sprintf("token:%d", id)
	if err := s.casbinService.RemoveSubjectPolicies(subject); err != nil {
		slog.Error("从Casbin移除令牌策略失败", "error", err)
	}

	if err := s.repo.DeleteTokenPermissions(s.ctx, id); err != nil {
		slog.Error("删除令牌权限关联失败", "error", err)
	}

	return s.repo.Delete(s.ctx, id)
}

func (s *ServiceTokenService) GetByID(id uint) (*modelSysTool.ServiceTokenDetailResp, error) {
	token, err := s.repo.FindByID(s.ctx, id)
	if err != nil {
		return nil, err
	}

	permissionIDs, err := s.repo.GetTokenPermissions(s.ctx, id)
	if err != nil {
		return nil, err
	}

	return &modelSysTool.ServiceTokenDetailResp{
		ServiceTokenResp: modelSysTool.ServiceTokenResp{
			ID:        token.ID,
			Name:      token.Name,
			Status:    token.Status,
			ExpiresAt: token.ExpiresAt,
			ApiCount:  len(permissionIDs),
			CreatedAt: token.CreatedAt.Unix(),
		},
		//ApiIDs: permissionIDs,
	}, nil
}

func (s *ServiceTokenService) List(req *modelSysTool.ListServiceTokenReq) (*modelSysTool.ServiceTokenListResp, error) {
	tokens, total, err := s.repo.List(s.ctx, req)
	if err != nil {
		return nil, err
	}

	list := make([]modelSysTool.ServiceTokenResp, 0, len(tokens))
	for _, token := range tokens {
		apiCount, _ := s.getTokenAPICount(token.ID)

		list = append(list, modelSysTool.ServiceTokenResp{
			ID:        token.ID,
			Name:      token.Name,
			Status:    token.Status,
			ExpiresAt: token.ExpiresAt,
			ApiCount:  apiCount,
			CreatedAt: token.CreatedAt.Unix(),
		})
	}

	return &modelSysTool.ServiceTokenListResp{
		List:  list,
		Total: total,
	}, nil
}

// AuthenticateToken 仅验证服务令牌是否有效（认证，不检查权限）
func (s *ServiceTokenService) AuthenticateToken(plainToken string) (tokenID uint, err error) {
	hash := sha256.Sum256([]byte(plainToken))
	tokenHash := hex.EncodeToString(hash[:])

	token, err := s.repo.FindByTokenHash(s.ctx, tokenHash)
	if err != nil {
		return 0, errors.New("invalid token")
	}

	if !token.Status {
		return 0, errors.New("token is disabled")
	}

	if token.ExpiresAt != nil && *token.ExpiresAt < time.Now().Unix() {
		return 0, errors.New("token is expired")
	}

	return token.ID, nil
}

// ValidateToken 使用Casbin验证权限（授权）
func (s *ServiceTokenService) ValidateToken(tokenID uint, path, method string) (bool, error) {
	subject := fmt.Sprintf("token:%d", tokenID)
	return s.casbinService.EnforceSubject(subject, path, method)
}

func (s *ServiceTokenService) getPermissionIDsByAPIs(apiIDs []uint) ([]uint, error) {
	var permissions []modelSysManagement.PermissionModel
	if err := global.TD27_DB.Where("domain_id IN ? AND domain = ?", apiIDs, modelSysManagement.PermissionDomainAPI).
		Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("get permissions by APIs failed: %w", err)
	}

	permissionIDs := make([]uint, 0, len(permissions))
	for _, p := range permissions {
		permissionIDs = append(permissionIDs, p.ID)
	}

	return permissionIDs, nil
}

func (s *ServiceTokenService) getTokenAPICount(tokenID uint) (int, error) {
	var count int64
	if err := global.TD27_DB.Model(&modelSysTool.TokenPermission{}).
		Where("token_id = ?", tokenID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s *ServiceTokenService) syncTokenToCasbin(tokenID uint, permissionIDs []uint) error {
	subject := fmt.Sprintf("token:%d", tokenID)

	var permissions []modelSysManagement.PermissionModel
	if err := global.TD27_DB.Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
		return err
	}

	slog.Info("syncTokenToCasbin",
		"tokenID", tokenID,
		"permissionCount", len(permissions),
		"permissionIDs", permissionIDs)

	policies := make([][]string, 0, len(permissions))
	for _, p := range permissions {
		policies = append(policies, []string{
			subject,
			p.Resource,
			string(p.Action),
		})
	}

	return s.casbinService.RebuildSubjectPolicies(subject, policies)
}
