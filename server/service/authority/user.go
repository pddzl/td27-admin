package authority

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"server/global"
	modelAuthority "server/model/authority"
	authorityReq "server/model/authority/request"
	authorityRes "server/model/authority/response"
	commonReq "server/model/common/request"
	"server/utils"
)

type UserService struct{}

func (us *UserService) GetUserInfo(userId uint) (userResults authorityRes.UserResult, err error) {
	err = global.TD27_DB.Table("authority_user").Select("authority_user.created_at,authority_user.id,authority_user.username,authority_user.phone,authority_user.email,authority_user.active,authority_user.role_model_id,authority_role.role_name").Joins("inner join authority_role on authority_user.role_model_id = authority_role.id").Where("authority_user.id = ?", userId).Scan(&userResults).Error
	return
}

// GetUsers 获取所有用户
func (us *UserService) GetUsers(pageInfo commonReq.PageInfo) ([]authorityRes.UserResult, int64, error) {
	var userResults []authorityRes.UserResult
	var total int64

	db := global.TD27_DB.Model(&modelAuthority.UserModel{})

	// 分页
	err := db.Count(&total).Error
	if err != nil {
		return userResults, total, fmt.Errorf("分页count err: %v", err)
	} else {
		limit := pageInfo.PageSize
		offset := pageInfo.PageSize * (pageInfo.Page - 1)
		db = db.Limit(limit).Offset(offset)
		// 左连接 查询出role_name
		db.Select("authority_user.id,authority_user.username,authority_user.phone,authority_user.email,authority_user.active,authority_user.role_model_id,authority_role.role_name").Joins("left join authority_role on authority_user.role_model_id = authority_role.id").Scan(&userResults)
	}

	return userResults, total, err
}

// DeleteUser 删除用户
func (us *UserService) DeleteUser(id uint) (err error) {
	return global.TD27_DB.Where("id = ?", id).Unscoped().Delete(&modelAuthority.UserModel{}).Error
}

// AddUser 添加用户
func (us *UserService) AddUser(instance *authorityReq.AddUser) (err error) {
	if errors.Is(global.TD27_DB.Where("id = ?", instance.RoleModelID).First(&modelAuthority.RoleModel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("角色不存在")
	}

	var userModel modelAuthority.UserModel
	userModel.Username = instance.Username
	userModel.Password = utils.MD5V([]byte(instance.Password))
	userModel.Phone = instance.Phone
	userModel.Email = instance.Email
	userModel.Active = instance.Active
	userModel.RoleModelID = instance.RoleModelID

	return global.TD27_DB.Create(&userModel).Error
}

// EditUser 编辑用户
func (us *UserService) EditUser(instance *authorityReq.EditUser) (*authorityRes.UserResult, error) {
	var userModel modelAuthority.UserModel
	// 用户是否存在
	if errors.Is(global.TD27_DB.Where("id = ?", instance.ID).First(&userModel).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("记录不存在")
	}

	// 角色是否存在
	var roleModel modelAuthority.RoleModel
	if errors.Is(global.TD27_DB.Where("id = ?", instance.RoleModelID).First(&roleModel).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("角色不存在")
	}

	err := global.TD27_DB.Model(&userModel).Updates(map[string]interface{}{"username": instance.Username, "phone": instance.Phone, "email": instance.Email, "active": instance.Active, "role_model_id": instance.RoleModelID}).Error

	if err != nil {
		return nil, err
	}

	var userResult authorityRes.UserResult

	userResult.ID = userModel.ID
	userResult.Username = userModel.Username
	userResult.Phone = userModel.Phone
	userResult.Email = userModel.Email
	userResult.Active = userModel.Active
	userResult.RoleName = roleModel.RoleName
	userResult.RoleModelID = userModel.RoleModelID

	return &userResult, nil
}

// ModifyPass 修改用户密码
func (us *UserService) ModifyPass(mp *authorityReq.ModifyPass) (err error) {
	var userModel modelAuthority.UserModel
	if errors.Is(global.TD27_DB.Where("id = ? and password = ?", mp.ID, utils.MD5V([]byte(mp.OldPassword))).First(&userModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("旧密码错误")
	}

	return global.TD27_DB.Model(&userModel).Update("password", utils.MD5V([]byte(mp.NewPassword))).Error
}

// SwitchActive 切换启用状态
func (us *UserService) SwitchActive(sa *authorityReq.SwitchActive) (err error) {
	var userModel modelAuthority.UserModel
	if errors.Is(global.TD27_DB.Where("id = ?", sa.ID).First(&userModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}

	return global.TD27_DB.Model(&userModel).Update("active", sa.Active).Error
}
