package router

import (
	"server/router/fileM"
	"server/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	FileM  fileM.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
