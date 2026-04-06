package sysManagement

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	"server/internal/model/sysManagement"
)

type RoleService struct {
	roleRepository sysManagement.RoleRepository
	userRepository sysManagement.UserRepository
	ctx            context.Context
}

func NewRoleService() *RoleService {
	return &RoleService{
		roleRepository: sysManagement.NewRoleRepo(global.TD27_DB),
		userRepository: sysManagement.NewUserEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (s *RoleService) List(req *common.PageInfo) ([]sysManagement.RoleModel, int64, error) {
	list, count, err := s.roleRepository.List(s.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (s *RoleService) Create(req *sysManagement.RoleModel) (*sysManagement.RoleModel, error) {
	create, err := s.roleRepository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}

	// 新角色创建时，自动赋予默认的公开API权限（登录、验证码等）
	// 这些权限在统一权限表中已存在，只需建立关联
	global.TD27_LOG.Info("创建角色成功，默认API权限已通过统一权限表管理", zap.Uint("roleId", create.ID))
	return create, err
}

func (s *RoleService) Delete(id uint) error {
	// check users exist with this role
	var count int64
	err := s.userRepository.CountUsersByRole(s.ctx, id, &count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该角色下面还有所属用户")
	}

	err = s.roleRepository.Delete(s.ctx, id)
	if err != nil {
		return err
	}

	// 清空menus关联
	//err = global.TD27_DB.Model(&roleModel).Association("Menus").Clear()
	err = s.roleRepository.DeleteRoleMenu(s.ctx, id)
	if err != nil {
		return fmt.Errorf("删除role关联menus err: %v", err)
	}

	// 删除角色的所有权限关联（通过统一权限表）
	if err = s.roleRepository.DeleteRoleMenu(s.ctx, id); err != nil {
		global.TD27_LOG.Warn("删除role权限关联失败", zap.Error(err))
	}

	// 重新加载Casbin策略
	go func() {
		if err = casbinService.ReloadPolicy(); err != nil {
			global.TD27_LOG.Error("重新加载Casbin策略失败", zap.Error(err))
		}
	}()

	return nil
}

func (s *RoleService) Update(req *sysManagement.UpdateRoleReq) error {
	err := s.roleRepository.Update(s.ctx, req)
	if err != nil {
		return err
	}
	return nil
}
