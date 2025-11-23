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

func (fr *FileRouter) InitFileRouter(rg *gin.RouterGroup) {
	base := rg.Group("file")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("upload", fr.fileApi.Upload)    // 文件上传
	record.GET("download", fr.fileApi.Download) // 下载文件
	record.GET("delete", fr.fileApi.Delete)     // 删除文件
	// without record
	base.POST("getFileList", fr.fileApi.GetFileList) // 分页获取文件信息
}
