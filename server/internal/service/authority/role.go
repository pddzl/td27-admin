package authority

import (
	"errors"
	"fmt"
	"server/internal/model/authority/menu"
	"server/internal/model/authority/role"
	authority2 "server/internal/model/authority/user"
	baseReq "server/internal/model/base/request"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"server/internal/global"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (rs *RoleService) List() ([]role.RoleModel, error) {
	var roleList []role.RoleModel
	err := global.TD27_DB.Preload("Menus").Find(&roleList).Error

	return roleList, err
}

func (rs *RoleService) Create(instance *role.RoleModel) (*role.RoleModel, error) {
	err := global.TD27_DB.Create(instance).Error
	if err == nil {
		if err = casbinService.EditCasbin(instance.ID, baseReq.DefaultCasbin()); err != nil {
			global.TD27_LOG.Error("更新casbin rule失败", zap.Error(err))
		}
	}
	return instance, err

}

func (rs *RoleService) Delete(id uint) (err error) {
	var roleModel role.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", id).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	if !errors.Is(global.TD27_DB.Where("role_model_id = ?", id).First(&authority2.UserModel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色下面还有所属用户")
	}

	err = global.TD27_DB.Unscoped().Delete(&roleModel).Error
	if err != nil {
		return fmt.Errorf("删除role err: %v", err)
	}

	// 清空menus关联
	err = global.TD27_DB.Model(&roleModel).Association("Menus").Clear()
	if err != nil {
		return fmt.Errorf("删除role关联menus err: %v", err)
	}

	// 删除对应casbin rule
	authorityId := strconv.Itoa(int(roleModel.ID))
	ok := casbinService.ClearCasbin(0, authorityId)
	if !ok {
		global.TD27_LOG.Warn("删除role关联casbin_rule失败")
	}
	return
}

func (rs *RoleService) Update(eRole role.EditRole) (err error) {
	var roleModel role.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", eRole.ID).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	return global.TD27_DB.Model(&roleModel).Update("role_name", eRole.RoleName).Error
}

// EditRoleMenu 编辑用户menu
func (rs *RoleService) EditRoleMenu(roleId uint, ids []uint) (err error) {
	var roleModel role.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", roleId).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	var menuModel []menu.MenuModel
	err = global.TD27_DB.Where("id in ?", ids).Find(&menuModel).Error
	if err != nil {
		global.TD27_LOG.Error("EditRoleMenu 查询menu", zap.Error(err))
		return err
	}

	err = global.TD27_DB.Model(&roleModel).Association("Menus").Replace(menuModel)
	if err != nil {
		global.TD27_LOG.Error("EditRoleMenu 替换menu", zap.Error(err))
		return err
	}

	return err
}
