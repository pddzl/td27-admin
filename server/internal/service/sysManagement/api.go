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
	ctx            context.Context
}

func NewApiService() *ApiService {
	return &ApiService{
		apiRepo:        modelSysManagement.NewApiRepo(global.TD27_DB),
		permissionRepo: modelSysManagement.NewPermissionRepo(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (s *ApiService) Create(req *modelSysManagement.CreateApiReq) (*modelSysManagement.ApiModel, error) {
	instance, err := s.apiRepo.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}

	return instance, err
}

func (s *ApiService) List(req *modelSysManagement.ListApiReq) ([]*modelSysManagement.ApiModel, int64, error) {
	list, count, err := s.apiRepo.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}

// ElTree 获取所有api tree
// element-plus el-tree的数据格式
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
	// 先获取API信息用于日志
	api, err := s.apiRepo.FindOne(s.ctx, id)
	if err != nil {
		return err
	}

	err = s.apiRepo.Delete(s.ctx, id)
	if err != nil {
		return err
	}

	// 重新加载Casbin策略
	go func() {
		if err = casbinService.ReloadPolicy(); err != nil {
			global.TD27_LOG.Error("重新加载Casbin策略失败", zap.Error(err))
		}
	}()

	global.TD27_LOG.Info("删除API",
		zap.String("path", api.Path),
		zap.String("method", api.Method))

	return nil
}

func (s *ApiService) DeleteByIds(ids []uint) error {
	err := s.apiRepo.DeleteByIds(s.ctx, ids)
	if err != nil {
		return err
	}

	// 重新加载Casbin策略
	go func() {
		if err = casbinService.ReloadPolicy(); err != nil {
			global.TD27_LOG.Error("重新加载Casbin策略失败", zap.Error(err))
		}
	}()

	return nil
}

func (s *ApiService) Update(req *modelSysManagement.UpdateApiReq) error {
	_, err := s.apiRepo.Update(s.ctx, req)
	if err != nil {
		return err
	}

	// 重新加载Casbin策略
	go func() {
		if err = casbinService.ReloadPolicy(); err != nil {
			global.TD27_LOG.Error("重新加载Casbin策略失败", zap.Error(err))
		}
	}()

	return nil
}

// UpdateRoleAPIPermissions 更新角色的API权限
func (s *ApiService) UpdateRoleAPIPermissions(roleId uint, apiIds []uint) error {
	return casbinService.UpdateRoleAPIPermissions(roleId, apiIds)
}
