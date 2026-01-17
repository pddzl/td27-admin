package sysManagement

import (
	"context"

	"server/internal/global"
	"server/internal/model/sysManagement"
)

type MenuService struct {
	menuRepository sysManagement.MenuRepository
	userRepository sysManagement.UserRepository
	ctx            context.Context
}

func NewMenuService() *MenuService {
	return &MenuService{
		menuRepository: sysManagement.NewMenuRepository(global.TD27_DB),
		userRepository: sysManagement.NewUserRepository(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (s *MenuService) List(userId uint) ([]*sysManagement.MenuModel, error) {
	var findOneUserReq sysManagement.FindOneUserReq
	findOneUserReq.ID = userId
	user, err := s.userRepository.FindOne(s.ctx, &findOneUserReq)
	if err != nil {
		return nil, err
	}

	list, err := s.menuRepository.List(s.ctx, user.RoleModelID)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *MenuService) Create(req *sysManagement.Menu) error {
	return s.menuRepository.Create(s.ctx, req)
}

func (s *MenuService) Update(req *sysManagement.UpdateMenuReq) error {
	return s.menuRepository.Update(s.ctx, req)
}

func (s *MenuService) Delete(id uint) error {
	return s.menuRepository.Delete(s.ctx, id)
}

// GetElTreeMenus 获取所有menu
func (s *MenuService) GetElTreeMenus(roleId uint) ([]*sysManagement.MenuModel, []uint, error) {
	menus, i, err := s.menuRepository.GetElTreeMenus(s.ctx, roleId)
	if err != nil {
		return nil, nil, err
	}

	return menus, i, err
}
