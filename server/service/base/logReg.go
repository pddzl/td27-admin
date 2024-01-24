package base

import (
	"errors"
	"server/global"
	modelAuthority "server/model/authority"
	"server/utils"
)

// 登录注册相关

type LogRegService struct{}

// Login 登陆校验
func (lr *LogRegService) Login(u *modelAuthority.UserModel) (userInter *modelAuthority.UserModel, err error) {
	var userModel modelAuthority.UserModel
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.TD27_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&userModel).Error
	if err != nil {
		return nil, errors.New("用户不存在或密码不正确")
	}
	if userModel.Active == false {
		return nil, errors.New("用户为禁用状态")
	}
	return &userModel, err
}
