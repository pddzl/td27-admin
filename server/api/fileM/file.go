package fileM

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/global"
	"server/model/common/response"
	fileMReq "server/model/fileM/request"
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

// GetFileList 分页获取文件信息
func (f *FileApi) GetFileList(c *gin.Context) {
	var params fileMReq.FileSearchParams
	_ = c.ShouldBindJSON(&params)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&params); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if list, total, err := fileService.GetFileList(params); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     params.Page,
			PageSize: params.PageSize,
		}, "获取成功", c)
	}
}
