package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type ApiRouter struct {
	apiApi *sysManagement.ApiApi
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{apiApi: sysManagement.NewApiApi()}
}

func (ur *ApiRouter) InitApiRouter(rg *gin.RouterGroup) {
	base := rg.Group("api")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("create", ur.apiApi.Create)
	record.POST("delete", ur.apiApi.Delete)
	record.POST("deleteByIds", ur.apiApi.DeleteByIds)
	record.POST("update", ur.apiApi.Update)
	record.POST("getElTree", ur.apiApi.GetElTree)
	// without record
	base.POST("list", ur.apiApi.List)
}
