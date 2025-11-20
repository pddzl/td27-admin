package sysSet

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/sysSet"
	"server/internal/middleware"
)

type DictDetailRouter struct {
	dictDetailApi *sysSet.DictDetailApi
}

func NewDictDetailRouter() *DictDetailRouter {
	return &DictDetailRouter{
		dictDetailApi: sysSet.NewDictDetailApi(),
	}
}

func (ddr *DictDetailRouter) InitDictDetailRouter(Router *gin.RouterGroup) {
	// record
	dictDetailRouter := Router.Group("dictDetail").Use(middleware.OperationRecord())
	dictDetailRouter.POST("delDictDetail", ddr.dictDetailApi.DelDictDetail)
	dictDetailRouter.POST("addDictDetail", ddr.dictDetailApi.AddDictDetail)
	dictDetailRouter.POST("editDictDetail", ddr.dictDetailApi.EditDictDetail)
	// not record
	dictDetailWithoutRouter := Router.Group("dictDetail")
	dictDetailWithoutRouter.POST("getDictDetail", ddr.dictDetailApi.GetDictDetail)
	dictDetailWithoutRouter.POST("getDictDetailFlat", ddr.dictDetailApi.GetDictDetailFlat)
}
