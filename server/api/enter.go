package api

import (
	"server/api/authority"
	"server/api/base"
	"server/api/fileM"
	"server/api/monitor"
	"server/api/sysSet"
	"server/api/sysTool"
)

type ApiGroup struct {
	Authority authority.ApiGroup
	Base      base.ApiGroup
	FileM     fileM.ApiGroup
	Monitor   monitor.ApiGroup
	SysTool   sysTool.ApiGroup
	SysSet    sysSet.SysSetGroup
}

var ApiGroupApp = new(ApiGroup)
