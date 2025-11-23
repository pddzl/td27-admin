package sysSet

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysSet"
	"server/internal/middleware"
)

type DictRouter struct {
	dictApi *sysSet.DictApi
}

func NewDictRouter() *DictRouter {
	return &DictRouter{
		dictApi: sysSet.NewDictApi(),
	}
}

func (dr *DictRouter) InitDictRouter(rg *gin.RouterGroup) {
	base := rg.Group("dict")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("delDict", dr.dictApi.DelDict)
	record.POST("addDict", dr.dictApi.AddDict)
	record.POST("editDict", dr.dictApi.EditDict)
	// not record
	base.GET("getDict", dr.dictApi.GetDict)
}
