package authority

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"

	"server/global"
	modelAuthority "server/model/authority"
	authorityReq "server/model/authority/request"
	baseReq "server/model/base/request"
	serviceBase "server/service/base"
)

type RoleService struct{}

func (rs *RoleService) GetRoles() ([]modelAuthority.RoleModel, error) {
	var roleList []modelAuthority.RoleModel
	err := global.TD27_DB.Preload("Menus").Find(&roleList).Error

	return roleList, err
}

func (rs *RoleService) AddRole(instance *modelAuthority.RoleModel) (*modelAuthority.RoleModel, error) {
	err := global.TD27_DB.Create(instance).Error
	if err == nil {
		if err = serviceBase.CasbinServiceApp.EditCasbin(instance.ID, baseReq.DefaultCasbin()); err != nil {
			global.TD27_LOG.Error("更新casbin rule失败", zap.Error(err))
		}
	}
	return instance, err

}

// DeleteRole 删除角色
func (rs *RoleService) DeleteRole(id uint) (err error) {
	var roleModel modelAuthority.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", id).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	if !errors.Is(global.TD27_DB.Where("role_model_id = ?", id).First(&modelAuthority.UserModel{}).Error, gorm.ErrRecordNotFound) {
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
	ok := serviceBase.CasbinServiceApp.ClearCasbin(0, authorityId)
	if !ok {
		global.TD27_LOG.Warn("删除role关联casbin_rule失败")
	}
	return
}

// EditRole 编辑用户
func (rs *RoleService) EditRole(eRole authorityReq.EditRole) (err error) {
	var roleModel modelAuthority.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", eRole.ID).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	return global.TD27_DB.Model(&roleModel).Update("role_name", eRole.RoleName).Error
}

// EditRoleMenu 编辑用户menu
func (rs *RoleService) EditRoleMenu(roleId uint, ids []uint) (err error) {
	var roleModel modelAuthority.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", roleId).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	var menuModel []modelAuthority.MenuModel
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
