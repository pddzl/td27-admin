package fileM

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/fileM"
	"server/internal/middleware"
)

type FileRouter struct {
	fileApi *fileM.FileApi
}

func NewFileRouter() *FileRouter {
	return &FileRouter{
		fileApi: fileM.NewFileApi(),
	}
}

func (fr *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	// record
	fileRouter := Router.Group("file").Use(middleware.OperationRecord())
	fileRouter.POST("upload", fr.fileApi.Upload)    // 文件上传
	fileRouter.GET("download", fr.fileApi.Download) // 下载文件
	fileRouter.GET("delete", fr.fileApi.Delete)     // 删除文件
	// without record
	fileWithoutRouter := Router.Group("file")
	fileWithoutRouter.POST("getFileList", fr.fileApi.GetFileList) // 分页获取文件信息
}
