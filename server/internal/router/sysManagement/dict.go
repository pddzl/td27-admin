package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type DictRouter struct {
	dictApi *sysManagement.DictApi
}

func NewDictRouter() *DictRouter {
	return &DictRouter{
		dictApi: sysManagement.NewDictApi(),
	}
}

func (r *DictRouter) InitDictRouter(rg *gin.RouterGroup) {
	base := rg.Group("dict")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("delete", r.dictApi.Delete)
	record.POST("create", r.dictApi.Create)
	record.POST("update", r.dictApi.Update)
	// not record
	base.POST("list", r.dictApi.List)
}
