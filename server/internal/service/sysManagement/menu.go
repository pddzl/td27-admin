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
		menuRepository: sysManagement.NewMenuEntity(global.TD27_DB),
		userRepository: sysManagement.NewUserEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (s *MenuService) List(userId uint) ([]sysManagement.MenuData, error) {
	// Get user with roles preloaded
	user, err := s.userRepository.FindOneWithRoles(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	// Get all role IDs from the user
	roleIDs := user.GetAllRoleIDs()
	if len(roleIDs) == 0 {
		return []sysManagement.MenuData{}, nil
	}

	// Get menus from all roles (union of permissions)
	items, err := s.menuRepository.ListByRoleIDs(s.ctx, roleIDs)
	if err != nil {
		return nil, err
	}

	// Convert to frontend format
	return sysManagement.ConvertToMenuDataList(items), nil
}

func (s *MenuService) Create(req *sysManagement.Menu) (*sysManagement.MenuModel, error) {
	return s.menuRepository.Create(s.ctx, req)
}

func (s *MenuService) Update(req *sysManagement.UpdateMenuReq) error {
	return s.menuRepository.Update(s.ctx, req)
}

func (s *MenuService) Delete(id uint) error {
	return s.menuRepository.Delete(s.ctx, id)
}

// GetElTreeMenus 获取所有menu
func (s *MenuService) GetElTreeMenus(roleId uint) ([]*sysManagement.MenuItem, []uint, error) {
	menus, checkedIds, err := s.menuRepository.GetElTreeMenus(s.ctx, roleId)
	if err != nil {
		return nil, nil, err
	}

	return menus, checkedIds, err
}
