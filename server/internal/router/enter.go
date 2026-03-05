package router

import (
	"github.com/gin-gonic/gin"

	"server/internal/router/sysManagement"
	"server/internal/router/sysMonitor"
	"server/internal/router/sysTool"
)

func NewSysMonitorRouterGroup() *sysMonitor.RouterGroup {
	return sysMonitor.NewRouterGroup()
}

func NewSysToolRouterGroup() *sysTool.RouterGroup {
	return sysTool.NewRouterGroup()
}

func NewSysManagementRouterGroup() *sysManagement.RouterGroup {
	return sysManagement.NewRouterGroup()
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

func RegisterAllModuleRouter() {
	Register(NewSysManagementRouterGroup())
	Register(NewSysMonitorRouterGroup())
	Register(NewSysToolRouterGroup())
}
