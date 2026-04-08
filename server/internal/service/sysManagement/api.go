package sysManagement

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/sysManagement"
)

type ApiService struct {
	repository sysManagement.APIRepository
	ctx        context.Context
}

func NewApiService() *ApiService {
	return &ApiService{
		repository: sysManagement.NewApiEntity(global.TD27_DB),
		ctx:        context.Background(),
	}
}

func (s *ApiService) Create(req *sysManagement.CreateApiReq) (*sysManagement.ApiModel, error) {
	instance, err := s.repository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}

	return instance, err
}

func (s *ApiService) List(req *sysManagement.ListApiReq) ([]*sysManagement.ApiModel, int64, error) {
	list, count, err := s.repository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}

// GetElTree 获取所有api tree
// element-plus el-tree的数据格式
func (s *ApiService) GetElTree(roleId uint) ([]*sysManagement.ApiTreeNode, []string, []uint, error) {
	list, err := s.repository.GetElTree(s.ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	// 前端 el-tree default-checked-keys
	checkedKey := make([]string, 0)
	checkedIds := make([]uint, 0)

	// 从统一权限表获取角色的API权限
	permissions, err := casbinService.GetRoleAPIPermissions(roleId)
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
	api, err := s.repository.FindOne(s.ctx, id)
	if err != nil {
		return err
	}

	err = s.repository.Delete(s.ctx, id)
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
	err := s.repository.DeleteByIds(s.ctx, ids)
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

func (s *ApiService) Update(req *sysManagement.UpdateApiReq) error {
	_, err := s.repository.Update(s.ctx, req)
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
