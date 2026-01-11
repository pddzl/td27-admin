package authority

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"server/internal/global"
	"server/internal/model/authority"
	"server/internal/model/common"
)

type RoleService struct {
	repository authority.RoleEntity
	ctx        context.Context
}

func NewRoleService() *RoleService {
	return &RoleService{
		repository: authority.NewDefaultRoleEntity(global.TD27_DB),
		ctx:        context.Background(),
	}
}

func (rs *RoleService) List(req *common.PageInfo) ([]authority.RoleModel, int64, error) {
	list, count, err := rs.repository.List(rs.ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (rs *RoleService) Create(req *authority.RoleModel) (*authority.RoleModel, error) {
	create, err := rs.repository.Create(rs.ctx, req)
	if err != nil {
		return nil, err
	}
	// todo
	// 更新casbin rule
	//if err = casbinService.EditCasbin(instance.ID, baseReq.DefaultCasbin()); err != nil {
	//	global.TD27_LOG.Error("更新casbin rule失败", zap.Error(err))
	//}
	return create, err

}

func (rs *RoleService) Delete(id uint) error {
	// todo
	// check users exist
	//if !errors.Is(global.TD27_DB.Where("role_model_id = ?", id).First(&authority2.UserModel{}).Error, gorm.ErrRecordNotFound) {
	//	return errors.New("该角色下面还有所属用户")
	//}

	err := rs.repository.Delete(rs.ctx, id)
	if err != nil {
		return err
	}

	// todo
	// 清空menus关联
	//err = global.TD27_DB.Model(&roleModel).Association("Menus").Clear()
	//if err != nil {
	//	return fmt.Errorf("删除role关联menus err: %v", err)
	//}

	// todo
	// 删除对应casbin rule
	//authorityId := strconv.Itoa(int(roleModel.ID))
	//ok := casbinService.ClearCasbin(0, authorityId)
	//if !ok {
	//	global.TD27_LOG.Warn("删除role关联casbin_rule失败")
	//}
	return nil
}

func (rs *RoleService) Update(req *authority.UpdateRoleReq) error {
	err := rs.repository.Update(rs.ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRoleMenu 编辑用户menu
func (rs *RoleService) UpdateRoleMenu(roleId uint, ids []uint) (err error) {
	var roleModel authority.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", roleId).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	// todo
	// search menu
	//var menuModel []menu.MenuModel
	//err = global.TD27_DB.Where("id in ?", ids).Find(&menuModel).Error
	//if err != nil {
	//	global.TD27_LOG.Error("EditRoleMenu 查询menu", zap.Error(err))
	//	return err
	//}

	//err = global.TD27_DB.Model(&roleModel).Association("Menus").Replace(menuModel)
	//if err != nil {
	//	global.TD27_LOG.Error("EditRoleMenu 替换menu", zap.Error(err))
	//	return err
	//}

	return err
}
