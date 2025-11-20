package router

import (
	"server/internal/router/authority"
	"server/internal/router/base"
	"server/internal/router/fileM"
	"server/internal/router/monitor"
	"server/internal/router/sysSet"
	"server/internal/router/sysTool"
)

type RouterGroup struct {
	Base      base.RouterGroup
	Authority authority.RouterGroup
	FileM     fileM.RouterGroup
}

func NewMonitorGroup() *monitor.RouterGroup {
	return monitor.NewRouterGroup()
}

func NewSysSetGroup() *sysSet.RouterGroup {
	return sysSet.NewRouterGroup()
}

func NewSysToolGroup() *sysTool.RouterGroup {
	return sysTool.NewRouterGroup()
}

var RouterGroupApp = new(RouterGroup)
