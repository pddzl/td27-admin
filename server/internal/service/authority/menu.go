package authority

import (
	"context"

	"server/internal/global"
	"server/internal/model/authority"
)

type MenuService struct {
	menuRepository authority.MenuEntity
	userRepository authority.UserEntity
	ctx            context.Context
}

func NewMenuService() *MenuService {
	return &MenuService{
		menuRepository: authority.NewDefaultMenuEntity(global.TD27_DB),
		userRepository: authority.NewDefaultUserEntity(global.TD27_DB),
		ctx:            context.Background(),
	}
}

func (ms *MenuService) List(userId uint) ([]*authority.MenuModel, error) {
	user, err := ms.userRepository.FindOne(ms.ctx, userId)
	if err != nil {
		return nil, err
	}

	list, err := ms.menuRepository.List(ms.ctx, user.RoleModelID)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (ms *MenuService) Create(req *authority.Menu) error {
	return ms.menuRepository.Create(ms.ctx, req)
}

func (ms *MenuService) Update(req *authority.UpdateMenuReq) error {
	return ms.menuRepository.Update(ms.ctx, req)
}

func (ms *MenuService) Delete(id uint) error {
	return ms.menuRepository.Delete(ms.ctx, id)
}

// GetElTreeMenus 获取所有menu
func (ms *MenuService) GetElTreeMenus(roleId uint) ([]*authority.MenuModel, []uint, error) {
	menus, i, err := ms.menuRepository.GetElTreeMenus(ms.ctx, roleId)
	if err != nil {
		return nil, nil, err
	}

	return menus, i, err
}
