package authority

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/authority"
	"server/internal/middleware"
)

type UserRouter struct {
	userApi *authority.UserApi
}

func NewUserRouter() *UserRouter {
	return &UserRouter{userApi: authority.NewUserApi()}
}

func (ur *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	base := rg.Group("user")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("delete", ur.userApi.Delete)
	record.POST("create", ur.userApi.Create)
	record.POST("update", ur.userApi.Update)
	record.POST("modifyPass", ur.userApi.ModifyPass)
	record.POST("switchActive", ur.userApi.SwitchActive)
	// without record
	base.GET("getUserInfo", ur.userApi.GetUserInfo)
	base.POST("list", ur.userApi.List)
}
