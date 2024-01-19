package authority

import (
	"server/service"
)

type ApiGroup struct {
	UserApi
	MenuApi
	RoleApi
	ApiApi
}

var (
	userService   = service.ServiceGroupApp.AuthorityServiceGroup.UserService
	menuService   = service.ServiceGroupApp.AuthorityServiceGroup.MenuService
	roleService   = service.ServiceGroupApp.AuthorityServiceGroup.RoleService
	apiService    = service.ServiceGroupApp.AuthorityServiceGroup.ApiService
	casbinService = service.ServiceGroupApp.BaseServiceGroup.CasbinService
)
