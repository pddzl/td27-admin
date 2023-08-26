package fileM

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
)

type FileApi struct{}

// Upload 上传文件
func (f *FileApi) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithStatusMessage(400, fmt.Sprintf("上传失败：%s", err.Error()), c)
		return
	}

	// 只允许上传csv文件
	//if file.Header.Get("Content-Type") != "text/csv" {
	//	response.FailWithStatusMessage(400, "只允许上传csv文件", c)
	//	return
	//}

	if fullPath, err := fileService.Upload(file); err != nil {
		response.FailWithStatusMessage(400, "上传失败", c)
		global.TD27_LOG.Error("上传失败", zap.Error(err))
	} else {
		response.OkWithDetailed(gin.H{"path": fullPath}, "上传成功", c)
	}
}
