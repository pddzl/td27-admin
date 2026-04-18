package sysTool

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"server/internal/global"
)

// ServiceTokenRepository 服务令牌仓库接口
type ServiceTokenRepository interface {
	Create(ctx context.Context, token *ServiceToken) error
	Update(ctx context.Context, token *ServiceToken) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*ServiceToken, error)
	FindByTokenHash(ctx context.Context, tokenHash string) (*ServiceToken, error)
	List(ctx context.Context, req *ListServiceTokenReq) ([]ServiceToken, int64, error)

	// Permission operations
	GetTokenPermissions(ctx context.Context, tokenID uint) ([]uint, error)
	SetTokenPermissions(ctx context.Context, tokenID uint, permissionIDs []uint) error
	DeleteTokenPermissions(ctx context.Context, tokenID uint) error
}

type serviceTokenRepo struct{}

func NewServiceTokenRepo(conn *gorm.DB) ServiceTokenRepository {
	return &serviceTokenRepo{}
}

func (r *serviceTokenRepo) Create(ctx context.Context, token *ServiceToken) error {
	if err := global.TD27_DB.WithContext(ctx).Create(token).Error; err != nil {
		return fmt.Errorf("create service token failed: %w", err)
	}
	return nil
}

func (r *serviceTokenRepo) Update(ctx context.Context, token *ServiceToken) error {
	if err := global.TD27_DB.WithContext(ctx).Save(token).Error; err != nil {
		return fmt.Errorf("update service token failed: %w", err)
	}
	return nil
}

func (r *serviceTokenRepo) Delete(ctx context.Context, id uint) error {
	result := global.TD27_DB.WithContext(ctx).Unscoped().Delete(&ServiceToken{}, id)
	if result.Error != nil {
		return fmt.Errorf("delete service token failed: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("service token not found")
	}
	return nil
}

func (r *serviceTokenRepo) FindByID(ctx context.Context, id uint) (*ServiceToken, error) {
	var token ServiceToken
	if err := global.TD27_DB.WithContext(ctx).First(&token, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service token not found")
		}
		return nil, fmt.Errorf("find service token failed: %w", err)
	}
	return &token, nil
}

func (r *serviceTokenRepo) FindByTokenHash(ctx context.Context, tokenHash string) (*ServiceToken, error) {
	var token ServiceToken
	if err := global.TD27_DB.WithContext(ctx).Where("token_hash = ?", tokenHash).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("service token not found")
		}
		return nil, fmt.Errorf("find service token by hash failed: %w", err)
	}
	return &token, nil
}

func (r *serviceTokenRepo) List(ctx context.Context, req *ListServiceTokenReq) ([]ServiceToken, int64, error) {
	var tokens []ServiceToken
	var total int64

	db := global.TD27_DB.WithContext(ctx).Model(&ServiceToken{})

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count service tokens failed: %w", err)
	}

	req.Normalize()

	if err := db.Order("created_at DESC").Limit(req.PageSize).Offset(req.Offset()).Find(&tokens).Error; err != nil {
		return nil, 0, fmt.Errorf("list service tokens failed: %w", err)
	}

	return tokens, total, nil
}

func (r *serviceTokenRepo) GetTokenPermissions(ctx context.Context, tokenID uint) ([]uint, error) {
	var permissionIDs []uint
	if err := global.TD27_DB.WithContext(ctx).
		Model(&TokenPermission{}).
		Where("token_id = ?", tokenID).
		Pluck("permission_id", &permissionIDs).Error; err != nil {
		return nil, fmt.Errorf("get token permissions failed: %w", err)
	}
	return permissionIDs, nil
}

func (r *serviceTokenRepo) SetTokenPermissions(ctx context.Context, tokenID uint, permissionIDs []uint) error {
	return global.TD27_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete existing permissions
		if err := tx.Where("token_id = ?", tokenID).Delete(&TokenPermission{}).Error; err != nil {
			return fmt.Errorf("delete old permissions failed: %w", err)
		}

		// Insert new permissions
		if len(permissionIDs) > 0 {
			records := make([]TokenPermission, 0, len(permissionIDs))
			for _, pid := range permissionIDs {
				records = append(records, TokenPermission{
					TokenID:      tokenID,
					PermissionID: pid,
				})
			}
			if err := tx.Create(&records).Error; err != nil {
				return fmt.Errorf("create permissions failed: %w", err)
			}
		}

		return nil
	})
}

func (r *serviceTokenRepo) DeleteTokenPermissions(ctx context.Context, tokenID uint) error {
	if err := global.TD27_DB.WithContext(ctx).Where("token_id = ?", tokenID).Delete(&TokenPermission{}).Error; err != nil {
		return fmt.Errorf("delete token permissions failed: %w", err)
	}
	return nil
}
