package authority

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userWithoutRouter := Router.Group("user")

	userApi := api.ApiGroupApp.Authority.UserApi
	{
		userRouter.POST("deleteUser", userApi.DeleteUser)
		userRouter.POST("addUser", userApi.AddUser)
		userRouter.POST("editUser", userApi.EditUser)
		userRouter.POST("modifyPass", userApi.ModifyPass)
		userRouter.POST("switchActive", userApi.SwitchActive)
	}
	{
		userWithoutRouter.GET("getUserInfo", userApi.GetUserInfo)
		userWithoutRouter.POST("getUsers", userApi.GetUsers)
	}
}
