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
	record.POST("addCron", cr.cronApi.AddCron)
	record.POST("deleteCron", cr.cronApi.DeleteCron)
	record.POST("deleteCronByIds", cr.cronApi.DeleteCronByIds)
	record.POST("editCron", cr.cronApi.EditCron)
	record.POST("switchOpen", cr.cronApi.SwitchOpen)
	// not record
	base.POST("getCronList", cr.cronApi.GetCronList)
}
