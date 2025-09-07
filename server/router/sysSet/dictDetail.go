package sysSet

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type DictDetailRouter struct{}

func (ddr *DictDetailRouter) InitDictDetailRouter(Router *gin.RouterGroup) {
	dictDetailRouter := Router.Group("dictDetail").Use(middleware.OperationRecord())
	dictDetailWithoutRouter := Router.Group("dictDetail")

	dictDetailApi := api.ApiGroupApp.SysSet.DictDetailApi
	{
		dictDetailRouter.POST("delDictDetail", dictDetailApi.DelDictDetail)
		dictDetailRouter.POST("addDictDetail", dictDetailApi.AddDictDetail)
		dictDetailRouter.POST("editDictDetail", dictDetailApi.EditDictDetail)
	}
	{
		dictDetailWithoutRouter.POST("getDictDetail", dictDetailApi.GetDictDetail)
	}
}
