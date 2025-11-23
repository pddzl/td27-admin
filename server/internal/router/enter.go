package router

import (
	"github.com/gin-gonic/gin"

	"server/internal/router/authority"
	"server/internal/router/base"
	"server/internal/router/fileM"
	"server/internal/router/monitor"
	"server/internal/router/sysSet"
	"server/internal/router/sysTool"
)

func NewBaseRouterGroup() *base.RouterGroup {
	return base.NewRouterGroup()
}

func NewMonitorRouterGroup() *monitor.RouterGroup {
	return monitor.NewRouterGroup()
}

func NewSysSetRouterGroup() *sysSet.RouterGroup {
	return sysSet.NewRouterGroup()
}

func NewSysToolRouterGroup() *sysTool.RouterGroup {
	return sysTool.NewRouterGroup()
}

func NewAuthorityRouterGroup() *authority.RouterGroup {
	return authority.NewRouterGroup()
}

func NewFileMRouterGroup() *fileM.RouterGroup {
	return fileM.NewRouterGroup()
}

type ModuleRouter interface {
	InitPublic(group *gin.RouterGroup)
	InitPrivate(group *gin.RouterGroup)
}

var modules []ModuleRouter

func Register(m ModuleRouter) {
	modules = append(modules, m)
}

func GetAllModules() []ModuleRouter {
	return modules
}

func init() {
	Register(NewAuthorityRouterGroup())
	Register(NewBaseRouterGroup())
	Register(NewFileMRouterGroup())
	Register(NewMonitorRouterGroup())
	Register(NewSysSetRouterGroup())
	Register(NewSysToolRouterGroup())
}
