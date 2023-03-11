package system

import (
	"server/service"
)

type ApiGroup struct {
	BaseApi
	UserApi
	MenuApi
	RoleApi
	ApiApi
	CasbinApi
	JwtApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService   = service.ServiceGroupApp.SystemServiceGroup.MenuService
	roleService   = service.ServiceGroupApp.SystemServiceGroup.RoleService
	apiService    = service.ServiceGroupApp.SystemServiceGroup.ApiService
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)
