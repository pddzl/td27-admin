package system

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
	"strconv"
)

type RoleService struct{}

func (rs *RoleService) GetRoles() ([]systemModel.RoleModel, error) {
	var roleList []systemModel.RoleModel
	err := global.TD27_DB.Preload("Menus").Find(&roleList).Error

	return roleList, err
}

func (rs *RoleService) AddRole(roleName string) (*systemModel.RoleModel, error) {
	var roleModel systemModel.RoleModel
	roleModel.RoleName = roleName
	return &roleModel, global.TD27_DB.Create(&roleModel).Error

}

// DeleteRole 删除角色
func (rs *RoleService) DeleteRole(id uint) (err error) {
	var roleModel systemModel.RoleModel

	err = global.TD27_DB.Where("id = ?", id).First(&roleModel).Error
	if err != nil {
		return fmt.Errorf("查询role -> %v", err)
	}

	if !errors.Is(global.TD27_DB.Where("role_model_id = ?", id).First(&systemModel.UserModel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色下面还有所属用户")
	}

	err = global.TD27_DB.Unscoped().Delete(&roleModel).Error
	if err != nil {
		return fmt.Errorf("删除role -> %v", err)
	}

	// 清空menus关联
	err = global.TD27_DB.Model(&roleModel).Association("Menus").Clear()
	if err != nil {
		return fmt.Errorf("删除role关联menus -> %v", err)
	}

	// 删除对应casbin rule
	authorityId := strconv.Itoa(int(roleModel.ID))
	ok := CasbinServiceApp.ClearCasbin(0, authorityId)
	if !ok {
		global.TD27_LOG.Warn("删除role关联casbin_rule失败")
	}
	return
}

// EditRole 编辑用户
func (rs *RoleService) EditRole(eRole systemReq.EditRole) (err error) {
	var roleModel systemModel.RoleModel
	err = global.TD27_DB.Where("id = ?", eRole.ID).First(&roleModel).Error
	if err != nil {
		global.TD27_LOG.Error("查询角色", zap.Error(err))
	}

	return global.TD27_DB.Model(&roleModel).Update("role_name", eRole.RoleName).Error
}

// EditRoleMenu 编辑用户menu
func (rs *RoleService) EditRoleMenu(roleId uint, ids []uint) (err error) {
	var menuModel []systemModel.MenuModel
	err = global.TD27_DB.Where("id in ?", ids).Find(&menuModel).Error
	if err != nil {
		global.TD27_LOG.Error("EditRoleMenu 查询menu", zap.Error(err))
		return err
	}

	var roleModel systemModel.RoleModel
	err = global.TD27_DB.Where("id = ?", roleId).First(&roleModel).Error
	if err != nil {
		global.TD27_LOG.Error("EditRoleMenu 查询role", zap.Error(err))
		return err
	}

	err = global.TD27_DB.Model(&roleModel).Association("Menus").Replace(menuModel)
	if err != nil {
		global.TD27_LOG.Error("EditRoleMenu 替换menu", zap.Error(err))
		return err
	}

	return err
}
