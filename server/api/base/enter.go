package base

import (
	"server/service"
)

type ApiGroup struct {
	LogRegApi
	CasbinApi
	JwtApi
}

var (
	jwtService    = service.ServiceGroupApp.BaseServiceGroup.JwtService
	logRegService = service.ServiceGroupApp.BaseServiceGroup.LogRegService
	casbinService = service.ServiceGroupApp.BaseServiceGroup.CasbinService
)
