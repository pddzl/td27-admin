package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type UserRouter struct {
	userApi *sysManagement.UserApi
}

func NewUserRouter() *UserRouter {
	return &UserRouter{userApi: sysManagement.NewUserApi()}
}

func (r *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	base := rg.Group("user")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("delete", r.userApi.Delete)
	record.POST("create", r.userApi.Create)
	record.POST("update", r.userApi.Update)
	record.POST("modify-passwd", r.userApi.ModifyPasswd)
	record.POST("switch-active", r.userApi.SwitchActive)
	// without record
	base.GET("get-user-info", r.userApi.GetUserInfo)
	base.POST("list", r.userApi.List)
}
