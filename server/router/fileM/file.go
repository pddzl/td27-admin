package fileM

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type FileRouter struct{}

func (f *FileRouter) InitFileRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	fileRouter := Router.Group("file").Use(middleware.OperationRecord())
	fileApi := api.ApiGroupApp.FileApiGroup.FileApi
	{
		fileRouter.POST("upload", fileApi.Upload)           // 文件上传
		fileRouter.POST("getFileList", fileApi.GetFileList) // 分页获取文件信息

	}
	return fileRouter
}
