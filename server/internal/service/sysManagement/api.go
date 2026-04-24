package sysManagement

import (
	"context"
	"fmt"
	"server/internal/pkg"


	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
	"log/slog"
)

type ApiService struct {
	apiRepo        modelSysManagement.APIRepository
	permissionRepo modelSysManagement.PermissionRepository
	casbinService  *CasbinService
	ctx            context.Context
}

func NewApiService() *ApiService {
	return &ApiService{
		apiRepo:        modelSysManagement.NewApiRepo(global.TD27_DB),
		permissionRepo: modelSysManagement.NewPermissionRepo(global.TD27_DB),
		casbinService:  NewCasbinService(),
		ctx:            context.Background(),
	}
}

func (s *ApiService) Create(req *modelSysManagement.CreateApiReq) (*modelSysManagement.ApiModel, error) {
	instance, err := s.apiRepo.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}

	permission := &modelSysManagement.PermissionModel{
		Name:     instance.Description,
		Domain:   modelSysManagement.PermissionDomainAPI,
		Resource: req.Path,
		Action:   modelSysManagement.HTTPMethodToAction(req.Method),
		DomainID: instance.ID,
	}

	if err = s.permissionRepo.Create(s.ctx, permission); err != nil {
		s.apiRepo.Delete(s.ctx, instance.ID)
		return nil, fmt.Errorf("create permission failed: %w", err)
	}

	return instance, nil
}

func (s *ApiService) List(req *modelSysManagement.ListApiReq) ([]*modelSysManagement.ApiModel, int64, error) {
	return s.apiRepo.List(s.ctx, req)
}

func (s *ApiService) ElTree(req *modelSysManagement.ApiTreeReq) ([]*modelSysManagement.ApiTreeNode, []uint, error) {
	list, domainIds, err := s.apiRepo.ElTree(s.ctx)
	if err != nil {
		return nil, nil, err
	}

	//checkedKey := make([]string, 0)
	checkedIds := make([]uint, 0)
	permissions := make([]modelSysManagement.PermissionModel, 0)

	if req.FromSource == "role" {
		permissions, err = s.permissionRepo.ListByRoleID(s.ctx, req.ID, modelSysManagement.PermissionDomainAPI)
	} else if req.FromSource == "token" {
		permissions, err = s.permissionRepo.ListByTokenID(s.ctx, req.ID, modelSysManagement.PermissionDomainAPI)
	} else {
		return nil, nil, fmt.Errorf("invalid fromSource")
	}

	if err != nil {
		slog.Error("获取角色API权限失败", "error", err)
	} else {
		for _, perm := range permissions {
			//checkedKey = append(checkedKey, fmt.Sprintf("%s,%s", perm.Resource, perm.Action))
			if pkg.IsContain(domainIds, perm.DomainID) {
				checkedIds = append(checkedIds, perm.DomainID)
			}
		}
	}

	return list, checkedIds, nil
}

func (s *ApiService) Delete(id uint) error {
	api, err := s.apiRepo.FindOne(s.ctx, id)
	if err != nil {
		return err
	}

	if err = s.permissionRepo.DeleteByDomainID(s.ctx, id, modelSysManagement.PermissionDomainAPI); err != nil {
		slog.Error("删除API权限失败", "error", err)
	}

	if err = s.apiRepo.Delete(s.ctx, id); err != nil {
		return err
	}

	action := modelSysManagement.HTTPMethodToAction(api.Method)
	if err = s.casbinService.RemoveResourcePolicy(api.Path, string(action)); err != nil {
		slog.Error("从Casbin移除API策略失败", "error", err)
	}

	return nil
}

func (s *ApiService) DeleteByIds(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	apis, err := s.apiRepo.FindByIds(s.ctx, ids)
	if err != nil {
		return err
	}

	for _, id := range ids {
		if err = s.permissionRepo.DeleteByDomainID(s.ctx, id, modelSysManagement.PermissionDomainAPI); err != nil {
			slog.Error("批量删除API权限失败", "apiId", id, "error", err)
		}
	}

	if err = s.apiRepo.DeleteByIds(s.ctx, ids); err != nil {
		return err
	}

	for _, api := range apis {
		action := modelSysManagement.HTTPMethodToAction(api.Method)
		if err = s.casbinService.RemoveResourcePolicy(api.Path, string(action)); err != nil {
			slog.Error("从Casbin批量移除API策略失败", "error", err)
		}
	}

	return nil
}

func (s *ApiService) Update(req *modelSysManagement.UpdateApiReq) (*modelSysManagement.ApiModel, error) {
	// 获取旧API信息用于Casbin清理
	oldAPI, err := s.apiRepo.FindOne(s.ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 更新API
	api, err := s.apiRepo.Update(s.ctx, req)
	if err != nil {
		return nil, err
	}

	// 查找并更新对应权限
	perm, err := s.permissionRepo.FindByDomainID(s.ctx, req.ID, modelSysManagement.PermissionDomainAPI)
	if err != nil {
		// 权限不存在则创建
		perm = &modelSysManagement.PermissionModel{
			Name:     req.Description,
			Domain:   modelSysManagement.PermissionDomainAPI,
			Resource: req.Path,
			Action:   modelSysManagement.HTTPMethodToAction(req.Method),
			DomainID: req.ID,
		}
		if createErr := s.permissionRepo.Create(s.ctx, perm); createErr != nil {
			slog.Error("创建API权限失败", "error", createErr)
		}
	} else {
		// 更新现有权限
		perm.Name = req.Description
		perm.Resource = req.Path
		perm.Action = modelSysManagement.HTTPMethodToAction(req.Method)
		if updateErr := s.permissionRepo.Update(s.ctx, perm); updateErr != nil {
			slog.Error("更新API权限失败", "error", updateErr)
		}
	}

	// 如果路径或方法改变，同步更新Casbin策略（角色 + 服务令牌）
	oldAction := modelSysManagement.HTTPMethodToAction(oldAPI.Method)
	newAction := modelSysManagement.HTTPMethodToAction(req.Method)
	if oldAPI.Path != req.Path || oldAction != newAction {
		if err = s.casbinService.UpdateResourcePolicies(oldAPI.Path, oldAction.String(), req.Path, newAction.String()); err != nil {
			slog.Error("更新API Casbin策略失败", "error", err)
		}
	}

	return api, nil
}
