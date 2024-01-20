package service

import (
	"server/service/authority"
	"server/service/base"
	"server/service/fileM"
	"server/service/monitor"
)

type ServiceGroup struct {
	Base      base.ServiceGroup
	Authority authority.ServiceGroup
	FileM     fileM.ServiceGroup
	Monitor   monitor.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
