package sysManagement

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
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
	// 1. 创建API
	instance, err := s.apiRepo.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}

	// 2. 创建对应的权限 (domain_id = api.id)
	permission := &modelSysManagement.PermissionModel{
		Name:     instance.Description,
		Domain:   modelSysManagement.PermissionDomainAPI,
		Resource: req.Path,
		Action:   modelSysManagement.HTTPMethodToAction(req.Method),
		DomainID: instance.ID,
	}

	if err = s.permissionRepo.Create(s.ctx, permission); err != nil {
		// 权限创建失败，删除已创建的API
		s.apiRepo.Delete(s.ctx, instance.ID)
		return nil, fmt.Errorf("create permission failed: %w", err)
	}

	return instance, nil
}

func (s *ApiService) List(req *modelSysManagement.ListApiReq) ([]*modelSysManagement.ApiModel, int64, error) {
	return s.apiRepo.List(s.ctx, req)
}

// ElTree 获取所有api tree
func (s *ApiService) ElTree(roleId uint) ([]*modelSysManagement.ApiTreeNode, []string, []uint, error) {
	list, err := s.apiRepo.ElTree(s.ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	// 前端 el-tree default-checked-keys
	checkedKey := make([]string, 0)
	checkedIds := make([]uint, 0)

	// 从统一权限表获取角色的API权限
	permissions, err := s.permissionRepo.List(s.ctx, roleId, modelSysManagement.PermissionDomainAPI)
	if err != nil {
		global.TD27_LOG.Error("获取角色API权限失败", zap.Error(err))
	} else {
		for _, perm := range permissions {
			checkedKey = append(checkedKey, fmt.Sprintf("%s,%s", perm.Resource, perm.Action))
			checkedIds = append(checkedIds, perm.ID)
		}
	}

	return list, checkedKey, checkedIds, nil
}

func (s *ApiService) Delete(id uint) error {
	// 1. 获取API信息（用于后续从Casbin中移除）
	api, err := s.apiRepo.FindOne(s.ctx, id)
	if err != nil {
		return err
	}

	// 2. 删除对应的权限 (通过domain_id关联)
	if err = s.permissionRepo.DeleteByDomainID(s.ctx, id, modelSysManagement.PermissionDomainAPI); err != nil {
		global.TD27_LOG.Error("删除API权限失败", zap.Error(err))
	}

	// 3. 删除API
	if err = s.apiRepo.Delete(s.ctx, id); err != nil {
		return err
	}

	// 4. 从Casbin中移除该API相关的所有策略
	// 由于无法确定哪些角色有这个权限，需要重新加载策略
	// 或者使用RemoveFilteredPolicy来移除所有涉及该resource的策略
	action := modelSysManagement.HTTPMethodToAction(api.Method)
	if err = s.casbinService.RemoveResourcePolicy(api.Path, string(action)); err != nil {
		global.TD27_LOG.Error("从Casbin移除API策略失败", zap.Error(err))
	}

	return nil
}

func (s *ApiService) DeleteByIds(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	// 获取所有要删除的API信息
	apis, err := s.apiRepo.FindByIds(s.ctx, ids)
	if err != nil {
		return err
	}

	// 删除这些API对应的权限
	for _, id := range ids {
		if err = s.permissionRepo.DeleteByDomainID(s.ctx, id, modelSysManagement.PermissionDomainAPI); err != nil {
			global.TD27_LOG.Error("批量删除API权限失败", zap.Uint("apiId", id), zap.Error(err))
		}
	}

	// 删除API
	if err = s.apiRepo.DeleteByIds(s.ctx, ids); err != nil {
		return err
	}

	// 从Casbin中移除策略
	for _, api := range apis {
		action := modelSysManagement.HTTPMethodToAction(api.Method)
		if err = s.casbinService.RemoveResourcePolicy(api.Path, string(action)); err != nil {
			global.TD27_LOG.Error("从Casbin批量移除API策略失败", zap.Error(err))
		}
	}

	return nil
}

func (s *ApiService) Update(req *modelSysManagement.UpdateApiReq) (*modelSysManagement.ApiModel, error) {
	return s.apiRepo.Update(s.ctx, req)
}
