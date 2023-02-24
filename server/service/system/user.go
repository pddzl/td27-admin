package system

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
)

type UserService struct{}

// Login 登陆校验
func (us *UserService) Login(u *systemModel.UserModel) (userInter *systemModel.UserModel, err error) {
	var userModel systemModel.UserModel
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.TD27_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&userModel).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if userModel.Active == false {
		return nil, errors.New("用户为禁用状态")
	}
	return &userModel, err
}

// GetUsers 获取所有用户
func (us *UserService) GetUsers(pageInfo request.PageInfo) ([]systemRes.UserResult, int64, error) {
	var userResults []systemRes.UserResult
	var total int64

	db := global.TD27_DB.Model(&systemModel.UserModel{})

	// 分页
	err := db.Count(&total).Error
	if err != nil {
		return userResults, total, fmt.Errorf("分页count -> %v", err)
	} else {
		limit := pageInfo.PageSize
		offset := pageInfo.PageSize * (pageInfo.Page - 1)
		db = db.Limit(limit).Offset(offset)
		//err = db.Find(&list).Error
		// 左连接 查询出role_name
		db.Select("sys_user.id,sys_user.username,sys_user.phone,sys_user.email,sys_user.active,sys_user.role_model_id,sys_role.role_name").Joins("left join sys_role on sys_user.role_model_id = sys_role.id").Scan(&userResults)
	}

	return userResults, total, err
}

// DeleteUser 删除用户
func (us *UserService) DeleteUser(id uint) (err error) {
	return global.TD27_DB.Where("id = ?", id).Unscoped().Delete(&systemModel.UserModel{}).Error
}

// AddUser 添加用户
func (us *UserService) AddUser(user systemReq.AddUser) (err error) {
	err = global.TD27_DB.Where("id = ?", user.RoleModelID).First(&systemModel.RoleModel{}).Error
	if err != nil {
		global.TD27_LOG.Error("添加用户 -> 查询role", zap.Error(err))
		return err
	}

	var userModel systemModel.UserModel
	userModel.Username = user.Username
	userModel.Password = utils.MD5V([]byte(user.Password))
	userModel.Phone = user.Phone
	userModel.Email = user.Email
	userModel.Active = user.Active
	userModel.RoleModelID = user.RoleModelID

	return global.TD27_DB.Create(&userModel).Error
}

// EditUser 编辑用户
func (us *UserService) EditUser(user systemReq.EditUser) (*systemRes.UserResult, error) {
	var userModel systemModel.UserModel
	var userResult systemRes.UserResult
	// 用户是否存在
	err := global.TD27_DB.Where("id = ?", user.Id).First(&userModel).Error
	if err != nil {
		global.TD27_LOG.Error("编辑用户 -> 查询Id", zap.Error(err))
		return nil, err
	}

	// 角色是否存在
	var roleModel systemModel.RoleModel
	err = global.TD27_DB.Where("id = ?", user.RoleModelID).First(&roleModel).Error
	if err != nil {
		global.TD27_LOG.Error("编辑用户 -> 查询role", zap.Error(err))
		return nil, err
	}

	updateV := make(map[string]interface{}, 5)
	updateV["username"] = user.Username
	updateV["active"] = user.Active
	updateV["role_model_id"] = user.RoleModelID
	updateV["phone"] = user.Phone
	updateV["email"] = user.Email

	err = global.TD27_DB.Model(&userModel).Updates(updateV).Error
	if err != nil {
		global.TD27_LOG.Error("编辑用户 -> update", zap.Error(err))
		return nil, err
	}

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
func (us *UserService) ModifyPass(mp systemReq.ModifyPass) (err error) {
	var userModel systemModel.UserModel
	err = global.TD27_DB.Where("id = ? and password = ?", mp.Id, utils.MD5V([]byte(mp.OldPassword))).First(&userModel).Error
	if err != nil {
		global.TD27_LOG.Error("修改用户密码 -> 查询", zap.Error(err))
		return err
	}
	return global.TD27_DB.Model(&userModel).Update("password", utils.MD5V([]byte(mp.NewPassword))).Error
}
