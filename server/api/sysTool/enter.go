package sysTool

import "server/service"

type ApiGroup struct {
	CronApi
}

var (
	cronService = service.ServiceGroupApp.SysTool.CronService
)
