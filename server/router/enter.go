package router

import (
	"server/router/authority"
	"server/router/base"
	"server/router/fileM"
	"server/router/monitor"
)

type RouterGroup struct {
	Base      base.RouterGroup
	Authority authority.RouterGroup
	FileM     fileM.RouterGroup
	Monitor   monitor.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
