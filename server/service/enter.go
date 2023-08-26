package service

import (
	"server/service/fileM"
	"server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	FileMServiceGroup  fileM.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
