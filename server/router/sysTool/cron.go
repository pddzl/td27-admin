package sysTool

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type CronRouter struct{}

func (c *CronRouter) InitCronRouter(Router *gin.RouterGroup) {
	cronRouter := Router.Group("cron").Use(middleware.OperationRecord())
	cronWithoutRouter := Router.Group("cron")

	cronApi := api.ApiGroupApp.SysTool.CronApi
	{
		cronRouter.POST("addCron", cronApi.AddCron)
		cronRouter.POST("deleteCron", cronApi.DeleteCron)
		cronRouter.POST("deleteCronByIds", cronApi.DeleteCronByIds)
		cronRouter.POST("editCron", cronApi.EditCron)
		cronRouter.POST("switchOpen", cronApi.SwitchOpen)
	}
	{
		cronWithoutRouter.POST("getCronList", cronApi.GetCronList)
	}
}
