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
func (us *UserService) GetUsers(pageInfo request.PageInfo) ([]systemModel.UserModel, int64, error) {
	var list []systemModel.UserModel
	var total int64

	db := global.TD27_DB.Model(&systemModel.UserModel{})

	// 分页
	err := db.Count(&total).Error
	if err != nil {
		return list, total, fmt.Errorf("分页count -> %v", err)
	} else {
		limit := pageInfo.PageSize
		offset := pageInfo.PageSize * (pageInfo.Page - 1)
		db = db.Limit(limit).Offset(offset)
		err = db.Preload("Roles").Find(&list).Error
	}

	return list, total, err
}
