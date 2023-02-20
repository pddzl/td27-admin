package system

import (
	"fmt"
	"server/global"
	"server/model/common/request"
	systemModel "server/model/system"
	"server/utils"
)

type UserService struct{}

// Login 登陆校验
func (us *UserService) Login(u *systemModel.UserModel) (userInter *systemModel.UserModel, err error) {
	var userModel systemModel.UserModel
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.TD27_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&userModel).Error
	return &userModel, err
}

// GetUsers 获取所有用户
func (us *UserService) GetUsers(pageInfo request.PageInfo) ([]systemModel.UserResult, int64, error) {
	var userResults []systemModel.UserResult
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
		db.Debug().Select("sys_user.id,sys_user.username,sys_user.phone,sys_user.email,sys_user.active,sys_user.role_model_id,sys_role.role_name").Joins("left join sys_role on sys_user.role_model_id = sys_role.id").Scan(&userResults)
	}

	return userResults, total, err
}

// DeleteUser 删除用户
func (us *UserService) DeleteUser(id uint) (err error) {
	return global.TD27_DB.Where("id = ?", id).Unscoped().Delete(&systemModel.UserModel{}).Error
}
