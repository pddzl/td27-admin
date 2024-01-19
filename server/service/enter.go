package service

import (
	"server/service/authority"
	"server/service/base"
	"server/service/fileM"
	"server/service/monitor"
)

type ServiceGroup struct {
	BaseServiceGroup      base.ServiceGroup
	AuthorityServiceGroup authority.ServiceGroup
	FileMServiceGroup     fileM.ServiceGroup
	MonitorServiceGroup   monitor.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
