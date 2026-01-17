package sysManagement

import (
	"errors"
	"fmt"

	"server/internal/global"
	"server/internal/model/sysManagement"
	"server/internal/pkg"
)

// 登录注册相关

type LogRegService struct{}

func NewLogRegService() *LogRegService {
	return &LogRegService{}
}

// Login 登陆校验
func (lr *LogRegService) Login(u *sysManagement.UserModel) (userInter *sysManagement.UserModel, err error) {
	var userModel sysManagement.UserModel
	u.Password = pkg.MD5V([]byte(u.Password))
	err = global.TD27_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&userModel).Error
	if err != nil {
		return nil, fmt.Errorf("usrname or password error: %s", err.Error())
	}
	if userModel.Active == false {
		return nil, errors.New("用户为禁用状态")
	}
	return &userModel, err
}
