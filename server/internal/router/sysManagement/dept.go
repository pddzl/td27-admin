package sysManagement

import (
	"github.com/gin-gonic/gin"

	apiSysManagement "server/internal/api/sysManagement"
	"server/internal/middleware"
)

type DeptRouter struct {
	deptApi *apiSysManagement.DeptApi
}

func NewDeptRouter() *DeptRouter {
	return &DeptRouter{
		deptApi: apiSysManagement.NewDeptApi(),
	}
}

func (d *DeptRouter) InitDeptRouter(Router *gin.RouterGroup) {
	deptRouter := Router.Group("dept").
		Use(middleware.OperationRecord()).
		Use(middleware.CasbinHandler())

	{
		deptRouter.POST("list", d.deptApi.List)
		deptRouter.POST("create", d.deptApi.Create)
		deptRouter.POST("update", d.deptApi.Update)
		deptRouter.POST("delete", d.deptApi.Delete)
		deptRouter.POST("getElTreeDepts", d.deptApi.GetElTreeDepts)
	}
}
