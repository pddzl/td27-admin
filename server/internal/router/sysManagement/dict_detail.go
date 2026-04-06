package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type DictDetailRouter struct {
	dictDetailApi *sysManagement.DictDetailApi
}

func NewDictDetailRouter() *DictDetailRouter {
	return &DictDetailRouter{
		dictDetailApi: sysManagement.NewDictDetailApi(),
	}
}

func (r *DictDetailRouter) InitDictDetailRouter(rg *gin.RouterGroup) {
	base := rg.Group("dictDetail")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("delete", r.dictDetailApi.Delete)
	record.POST("create", r.dictDetailApi.Create)
	record.POST("update", r.dictDetailApi.Update)
	// not record
	base.POST("list", r.dictDetailApi.List)
	base.POST("flat", r.dictDetailApi.Flat)
}
