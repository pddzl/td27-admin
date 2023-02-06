package system

import (
	"server/service"
)

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
)
