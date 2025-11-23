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

func (ddr *DictDetailRouter) InitDictDetailRouter(rg *gin.RouterGroup) {
	base := rg.Group("dictDetail")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("delDictDetail", ddr.dictDetailApi.DelDictDetail)
	record.POST("addDictDetail", ddr.dictDetailApi.AddDictDetail)
	record.POST("editDictDetail", ddr.dictDetailApi.EditDictDetail)
	// not record
	base.POST("getDictDetail", ddr.dictDetailApi.GetDictDetail)
	base.POST("getDictDetailFlat", ddr.dictDetailApi.GetDictDetailFlat)
}
