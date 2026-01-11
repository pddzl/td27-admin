package authority

import (
	"context"

	"server/internal/global"
	"server/internal/model/authority/menu"
)

type MenuService struct {
	repository menu.MenuEntity
	ctx        context.Context
}

func NewMenuService() *MenuService {
	return &MenuService{
		repository: menu.NewDefaultMenuEntity(global.TD27_DB),
		ctx:        context.Background(),
	}
}

func (ms *MenuService) List(userId uint) ([]*menu.MenuModel, error) {
	// todo 查找用户对应角色
	list, err := ms.repository.List(ms.ctx, userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (ms *MenuService) Create(req *menu.Menu) error {
	return ms.repository.Create(ms.ctx, req)
}

func (ms *MenuService) Update(req *menu.UpdateMenuReq) error {
	return ms.repository.Update(ms.ctx, req)
}

func (ms *MenuService) Delete(id uint) error {
	return ms.repository.Delete(ms.ctx, id)
}

// GetElTreeMenus 获取所有menu
func (ms *MenuService) GetElTreeMenus(roleId uint) ([]*menu.MenuModel, []uint, error) {
	menus, i, err := ms.repository.GetElTreeMenus(ms.ctx, roleId)
	if err != nil {
		return nil, nil, err
	}

	return menus, i, err
}
