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

func (s *MenuService) List(userId uint) ([]*authority.MenuModel, error) {
	var findOneUserReq authority.FindOneUserReq
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

func (s *MenuService) Create(req *authority.Menu) error {
	return s.menuRepository.Create(s.ctx, req)
}

func (s *MenuService) Update(req *authority.UpdateMenuReq) error {
	return s.menuRepository.Update(s.ctx, req)
}

func (s *MenuService) Delete(id uint) error {
	return s.menuRepository.Delete(s.ctx, id)
}

// GetElTreeMenus 获取所有menu
func (s *MenuService) GetElTreeMenus(roleId uint) ([]*authority.MenuModel, []uint, error) {
	menus, i, err := s.menuRepository.GetElTreeMenus(s.ctx, roleId)
	if err != nil {
		return nil, nil, err
	}

	return menus, i, err
}
