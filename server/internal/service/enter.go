package service

import (
	"server/internal/service/authority"
	"server/internal/service/base"
	"server/internal/service/fileM"
	"server/internal/service/monitor"
	"server/internal/service/sysSet"
	"server/internal/service/sysTool"
)

type ServiceGroup struct {
	Base      base.ServiceGroup
	Authority authority.ServiceGroup
	FileM     fileM.ServiceGroup
	Monitor   monitor.ServiceGroup
	SysTool   sysTool.ServiceGroup
	SysSet    sysSet.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
