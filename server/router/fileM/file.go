package fileM

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type FileRouter struct{}

func (f *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file").Use(middleware.OperationRecord())
	fileWithoutRouter := Router.Group("file")

	fileApi := api.ApiGroupApp.FileM.FileApi
	{
		fileRouter.POST("upload", fileApi.Upload)    // 文件上传
		fileRouter.GET("download", fileApi.Download) // 下载文件
		fileRouter.GET("delete", fileApi.Delete)     // 删除文件
	}
	{
		fileWithoutRouter.POST("getFileList", fileApi.GetFileList) // 分页获取文件信息
	}
}
