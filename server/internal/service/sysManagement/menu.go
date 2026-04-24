package sysManagement

import (
	"context"
	"fmt"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
)

type MenuService struct {
	menuRepository modelSysManagement.MenuRepository
	userRepository modelSysManagement.UserRepository
	permissionRepo modelSysManagement.PermissionRepository
	casbinService  *CasbinService
	ctx            context.Context
}

func NewMenuService() *MenuService {
	return &MenuService{
		menuRepository: modelSysManagement.NewMenuRepo(global.TD27_DB),
		userRepository: modelSysManagement.NewUserEntity(global.TD27_DB),
		permissionRepo: modelSysManagement.NewPermissionRepo(global.TD27_DB),
		casbinService:  NewCasbinService(),
		ctx:            context.Background(),
	}
}

func (s *MenuService) List(customClaims *modelSysManagement.CustomClaims) ([]modelSysManagement.MenuResp, error) {
	roleIDs := customClaims.GetAllRoleIDs()
	return s.menuRepository.List(s.ctx, roleIDs)
}

func (s *MenuService) Create(req *modelSysManagement.Menu) (*modelSysManagement.MenuModel, error) {
	// 1. 创建菜单
	menu, err := s.menuRepository.Create(s.ctx, req)
	if err != nil {
		return nil, err
	}

	// 2. 创建对应的权限 (domain_id = menu.id)
	permission := &modelSysManagement.PermissionModel{
		Name:     req.Title,
		Domain:   modelSysManagement.PermissionDomainMenu,
		Resource: req.Path,
		Action:   modelSysManagement.ActionView,
		DomainID: menu.ID,
	}

	if err = s.permissionRepo.Create(s.ctx, permission); err != nil {
		global.TD27_LOG.Info(fmt.Sprintf("权限创建失败，删除已创建的菜单"))
		errDel := s.menuRepository.Delete(s.ctx, menu.ID)
		if errDel != nil {
			global.TD27_LOG.Error(fmt.Sprintf("删除菜单失败 menu_id %d, error: %v", menu.ID, errDel))
		}
		return nil, fmt.Errorf("create permission failed: %w", err)
	}

	return menu, nil
}

func (s *MenuService) Update(req *modelSysManagement.UpdateMenuReq) error {
	// 更新菜单
	if err := s.menuRepository.Update(s.ctx, req); err != nil {
		return err
	}

	// 查找并更新对应权限
	perm, err := s.permissionRepo.FindByDomainID(s.ctx, req.ID, modelSysManagement.PermissionDomainMenu)
	if err != nil {
		// 权限不存在则创建
		perm = &modelSysManagement.PermissionModel{
			Name:     req.Title,
			Domain:   modelSysManagement.PermissionDomainMenu,
			Resource: req.Path,
			Action:   modelSysManagement.ActionView,
			DomainID: req.ID,
		}
		if createErr := s.permissionRepo.Create(s.ctx, perm); createErr != nil {
			global.TD27_LOG.Error("创建菜单权限失败", "error", createErr)
		}
	} else {
		// 更新现有权限
		perm.Name = req.Title
		perm.Resource = req.Path
		if updateErr := s.permissionRepo.Update(s.ctx, perm); updateErr != nil {
			global.TD27_LOG.Error("更新菜单权限失败", "error", updateErr)
		}
	}

	return nil
}

func (s *MenuService) Delete(id uint) error {
	// 删除对应的权限 (通过domain_id关联)
	if err := s.permissionRepo.DeleteByDomainID(s.ctx, id, modelSysManagement.PermissionDomainMenu); err != nil {
		global.TD27_LOG.Error("删除菜单权限失败", "error", err)
	}

	return s.menuRepository.Delete(s.ctx, id)
}

// ElTree 获取所有menu
func (s *MenuService) ElTree(roleId uint) ([]modelSysManagement.MenuResp, []uint, error) {
	return s.menuRepository.ElTree(s.ctx, roleId)
}
