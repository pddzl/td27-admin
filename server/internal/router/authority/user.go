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
	record.POST("deleteUser", ur.userApi.DeleteUser)
	record.POST("addUser", ur.userApi.AddUser)
	record.POST("editUser", ur.userApi.EditUser)
	record.POST("modifyPass", ur.userApi.ModifyPass)
	record.POST("switchActive", ur.userApi.SwitchActive)
	// without record
	base.GET("getUserInfo", ur.userApi.GetUserInfo)
	base.POST("getUsers", ur.userApi.GetUsers)
}
