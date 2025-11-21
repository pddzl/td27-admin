package authority

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/authority"
	"server/internal/middleware"
)

type ApiRouter struct {
	apiApi *authority.ApiApi
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{apiApi: authority.NewApiApi()}
}

func (ur *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	// record
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouter.POST("addApi", ur.apiApi.AddApi)
	apiRouter.POST("deleteApi", ur.apiApi.DeleteApi)
	apiRouter.POST("deleteApiById", ur.apiApi.DeleteApiById)
	apiRouter.POST("editApi", ur.apiApi.EditApi)
	apiRouter.POST("getElTreeApis", ur.apiApi.GetElTreeApis)
	// without record
	apiWithoutRouter := Router.Group("api")
	apiWithoutRouter.POST("getApis", ur.apiApi.GetApis)
}
