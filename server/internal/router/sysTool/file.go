package sysTool

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysTool"
	"server/internal/middleware"
)

type FileRouter struct {
	fileApi *sysTool.FileApi
}

func NewFileRouter() *FileRouter {
	return &FileRouter{
		fileApi: sysTool.NewFileApi(),
	}
}

func (r *FileRouter) InitFileRouter(rg *gin.RouterGroup) {
	base := rg.Group("file")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("upload", r.fileApi.Upload)    // 文件上传
	record.GET("download", r.fileApi.Download) // 下载文件
	record.GET("delete", r.fileApi.Delete)     // 删除文件
	// without record
	base.POST("list", r.fileApi.List) // 分页获取文件信息
}
