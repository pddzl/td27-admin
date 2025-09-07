package sysSet

import "server/service"

type SysSetGroup struct {
	DictApi
	DictDetailApi
}

var (
	dictService       = service.ServiceGroupApp.SysSet.DictService
	dictDetailService = service.ServiceGroupApp.SysSet.DictDetailService
)
