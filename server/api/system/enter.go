package system

import (
	"server/service"
)

type ApiGroup struct {
	BaseApi
	UserApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
