package sysTool

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysTool"
	"server/internal/middleware"
)

type CronRouter struct {
	cronApi *sysTool.CronApi
}

func NewCronRouter() *CronRouter {
	return &CronRouter{
		cronApi: sysTool.NewCronApi(),
	}
}

func (cr *CronRouter) InitCronRouter(rg *gin.RouterGroup) {
	base := rg.Group("cron")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("create", cr.cronApi.Create)
	record.POST("delete", cr.cronApi.Delete)
	record.POST("deleteByIds", cr.cronApi.DeleteByIds)
	record.POST("update", cr.cronApi.Update)
	record.POST("switchOpen", cr.cronApi.SwitchOpen)
	// not record
	base.POST("list", cr.cronApi.List)
}
