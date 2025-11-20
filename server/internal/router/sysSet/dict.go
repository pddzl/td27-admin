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

func (dr *DictRouter) InitDictRouter(Router *gin.RouterGroup) {
	// record
	dictRouter := Router.Group("dict").Use(middleware.OperationRecord())
	dictRouter.POST("delDict", dr.dictApi.DelDict)
	dictRouter.POST("addDict", dr.dictApi.AddDict)
	dictRouter.POST("editDict", dr.dictApi.EditDict)
	// not record
	dictWithoutRouter := Router.Group("dict")
	dictWithoutRouter.GET("getDict", dr.dictApi.GetDict)

}
