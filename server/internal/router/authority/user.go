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

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// record
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouter.POST("deleteUser", ur.userApi.DeleteUser)
	userRouter.POST("addUser", ur.userApi.AddUser)
	userRouter.POST("editUser", ur.userApi.EditUser)
	userRouter.POST("modifyPass", ur.userApi.ModifyPass)
	userRouter.POST("switchActive", ur.userApi.SwitchActive)
	// without record
	userWithoutRouter := Router.Group("user")
	userWithoutRouter.GET("getUserInfo", ur.userApi.GetUserInfo)
	userWithoutRouter.POST("getUsers", ur.userApi.GetUsers)
}
