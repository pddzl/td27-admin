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

func (d *DeptRouter) InitDeptRouter(rg *gin.RouterGroup) {
	base := rg.Group("dept")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("create", d.deptApi.Create)
	record.POST("delete", d.deptApi.Delete)
	record.POST("update", d.deptApi.Update)
	record.POST("getElTreeDepts", d.deptApi.GetElTreeDepts)
	// not record
	base.POST("list", d.deptApi.List)
}
