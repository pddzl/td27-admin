package sysSet

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type DictRouter struct{}

func (dr *DictRouter) InitDictRouter(Router *gin.RouterGroup) {
	dictRouter := Router.Group("dict").Use(middleware.OperationRecord())
	dictWithoutRouter := Router.Group("dict")

	dictApi := api.ApiGroupApp.SysSet.DictApi
	{
		dictRouter.POST("delDict", dictApi.DelDict)
		dictRouter.POST("addDict", dictApi.AddDict)
		dictRouter.POST("editDict", dictApi.EditDict)
	}
	{
		dictWithoutRouter.GET("getDict", dictApi.GetDict)
	}
}
