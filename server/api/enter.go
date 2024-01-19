package api

import (
	"server/api/authority"
	"server/api/base"
	"server/api/fileM"
	"server/api/monitor"
)

type ApiGroup struct {
	AuthorityApiGroup authority.ApiGroup
	BaseApiGroup      base.ApiGroup
	FileMApiGroup     fileM.ApiGroup
	MonitorApiGroup   monitor.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
