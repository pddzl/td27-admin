package api

import (
	"server/api/authority"
	"server/api/base"
	"server/api/fileM"
	"server/api/monitor"
)

type ApiGroup struct {
	Authority authority.ApiGroup
	Base      base.ApiGroup
	FileM     fileM.ApiGroup
	Monitor   monitor.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
