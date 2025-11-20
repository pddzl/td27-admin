package api

import (
	"server/internal/api/authority"
	"server/internal/api/base"
	"server/internal/api/fileM"
	"server/internal/api/monitor"
	"server/internal/api/sysSet"
	"server/internal/api/sysTool"
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
