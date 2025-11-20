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

func (cr *CronRouter) InitCronRouter(Router *gin.RouterGroup) {
	// record
	cronRouter := Router.Group("cron").Use(middleware.OperationRecord())
	cronRouter.POST("addCron", cr.cronApi.AddCron)
	cronRouter.POST("deleteCron", cr.cronApi.DeleteCron)
	cronRouter.POST("deleteCronByIds", cr.cronApi.DeleteCronByIds)
	cronRouter.POST("editCron", cr.cronApi.EditCron)
	cronRouter.POST("switchOpen", cr.cronApi.SwitchOpen)
	// not record
	cronWithoutRouter := Router.Group("cron")
	cronWithoutRouter.POST("getCronList", cr.cronApi.GetCronList)
}
