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
	userService   = service.ServiceGroupApp.Authority.UserService
	menuService   = service.ServiceGroupApp.Authority.MenuService
	roleService   = service.ServiceGroupApp.Authority.RoleService
	apiService    = service.ServiceGroupApp.Authority.ApiService
	casbinService = service.ServiceGroupApp.Base.CasbinService
)
