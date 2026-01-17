package sysManagement

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	"server/internal/model/sysManagement"
)

type RoleService struct {
	roleRepository sysManagement.RoleRepository
	userRepository sysManagement.UserRepository
	menuRepository sysManagement.MenuRepository
	ctx            context.Context
}

func NewRoleService() *RoleService {
	return &RoleService{
		roleRepository: sysManagement.NewRoleRepository(global.TD27_DB),
		userRepository: sysManagement.NewUserRepository(global.TD27_DB),
		menuRepository: sysManagement.NewMenuRepository(global.TD27_DB),
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

	// 更新casbin rule
	if err = casbinService.EditCasbin(create.ID, sysManagement.DefaultCasbin()); err != nil {
		global.TD27_LOG.Error("更新casbin rule失败", zap.Error(err))
	}
	return create, err

}

func (s *RoleService) Delete(id uint) error {
	// check users exist
	var userReq sysManagement.FindOneUserReq
	userReq.RoleModelID = id
	user, err := s.userRepository.FindOne(s.ctx, &userReq)
	if err != nil {
		return err
	}
	if user != nil {
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

	// 删除对应casbin rule
	authorityId := strconv.Itoa(int(id))
	ok := casbinService.ClearCasbin(0, authorityId)
	if !ok {
		global.TD27_LOG.Warn("删除role关联casbin_rule失败")
	}

	return nil
}

func (s *RoleService) Update(req *sysManagement.UpdateRoleReq) error {
	err := s.roleRepository.Update(s.ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRoleMenu 编辑用户menu
func (s *RoleService) UpdateRoleMenu(roleId uint, menuIds []uint) error {
	// check role existence
	_, err := s.roleRepository.FindOne(s.ctx, roleId)
	if err != nil {
		return err
	}

	// search menu
	menus, err := s.menuRepository.FindByIds(s.ctx, menuIds)
	if err != nil {
		return fmt.Errorf("FindByIds err: %v", err)
	}

	// update role_menus
	err = s.roleRepository.UpdateRoleMenu(s.ctx, menus)
	if err != nil {
		return err
	}

	return err
}
