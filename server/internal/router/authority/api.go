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

func (ur *ApiRouter) InitApiRouter(rg *gin.RouterGroup) {
	base := rg.Group("api")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("addApi", ur.apiApi.AddApi)
	record.POST("deleteApi", ur.apiApi.DeleteApi)
	record.POST("deleteApiById", ur.apiApi.DeleteApiById)
	record.POST("editApi", ur.apiApi.EditApi)
	record.POST("getElTreeApis", ur.apiApi.GetElTreeApis)
	// without record
	base.POST("getApis", ur.apiApi.GetApis)
}
