package system

import (
	"server/global"
	"server/model/system"
	"server/utils"
)

type UserService struct{}

// Login 登陆校验
func (us *UserService) Login(u *system.UserModel) (userInter *system.UserModel, err error) {
	var userModel system.UserModel
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.TD27_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&userModel).Error
	return &userModel, err
}
