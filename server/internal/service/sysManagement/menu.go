package sysManagement

import (
	"context"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
)

type MenuService struct {
	menuRepository modelSysManagement.MenuRepository
	userRepository modelSysManagement.UserRepository
	ctx            context.Context
}

func NewMenuService() *MenuService {
	return &MenuService{
		menuRepository: modelSysManagement.NewMenuRepo(global.TD27_DB),
		userRepository: modelSysManagement.NewUserEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (s *MenuService) List(customClaims *modelSysManagement.CustomClaims) ([]modelSysManagement.MenuResp, error) {
	// Get all role IDs from the user
	roleIDs := customClaims.GetAllRoleIDs()

	// Get menus from all roles (union of permissions)
	return s.menuRepository.List(s.ctx, roleIDs)
}

func (s *MenuService) Create(req *modelSysManagement.Menu) (*modelSysManagement.MenuModel, error) {
	return s.menuRepository.Create(s.ctx, req)
}

func (s *MenuService) Update(req *modelSysManagement.UpdateMenuReq) error {
	return s.menuRepository.Update(s.ctx, req)
}

func (s *MenuService) Delete(id uint) error {
	return s.menuRepository.Delete(s.ctx, id)
}

// ElTree 获取所有menu
func (s *MenuService) ElTree(roleId uint) ([]modelSysManagement.MenuResp, []uint, error) {
	return s.menuRepository.ElTree(s.ctx, roleId)
}
