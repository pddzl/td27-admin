package fileM

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"

	"server/internal/global"
	"server/internal/model/common/response"
	fileMReq "server/internal/model/entity/fileM/request"
	"server/internal/service/fileM"
)

type FileApi struct {
	fileService *fileM.FileService
}

func NewFileApi() *FileApi {
	return &FileApi{
		fileService: fileM.NewFileService(),
	}
}

// Upload
// @Tags      FileApi
// @Summary   上传文件
// @Security  ApiKeyAuth
// @accept    mpfd
// @Produce   application/json
// @Param     file formData file true "The file to upload"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /file/upload [post]
func (fa *FileApi) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithStatusMessage(400, fmt.Sprintf("上传失败：%s", err.Error()), c)
		return
	}

	// 只允许上传csv文件
	if file.Header.Get("Content-Type") != "text/csv" {
		response.FailWithStatusMessage(400, "只允许上传csv文件", c)
		return
	}

	if fullInfo, err := fa.fileService.Upload(file); err != nil {
		response.FailWithStatusMessage(400, "上传失败", c)
		global.TD27_LOG.Error("上传失败", zap.Error(err))
	} else {
		response.OkWithDetailed(fullInfo, "上传成功", c)
	}
}

// GetFileList
// @Tags      FileApi
// @Summary   分页获取文件信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      fileMReq.FileSearchParams true  "请求参数"
// @Success   200   {object}  response.Response{data=response.Page{list=[]fileM.FileModel},msg=string}
// @Router    /file/getFileList [post]
func (fa *FileApi) GetFileList(c *gin.Context) {
	var params fileMReq.FileSearchParams
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := fa.fileService.GetFileList(params); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.Page{
			List:     list,
			Total:    total,
			Page:     params.Page,
			PageSize: params.PageSize,
		}, "获取成功", c)
	}
}

// Download
// @Tags      FileApi
// @Summary   下载文件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   octet-stream
// @Param     name query string true "文件名"
// @Success   200   {file} application/octet-stream
// @Router    /file/download [get]
func (fa *FileApi) Download(c *gin.Context) {
	fileName := c.Query("name")

	path := fmt.Sprintf("%s/%s", global.TD27_CONFIG.File.Upload, fileName)

	// 打开文件
	_, err := os.Stat(path)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("文件错误：%v", err), c)
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(path)
}

// Delete
// @Tags      FileApi
// @Summary   删除文件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     name query string true "文件名"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /file/delete [get]
func (fa *FileApi) Delete(c *gin.Context) {
	fileName := c.Query("name")

	if err := fa.fileService.Delete(fileName); err != nil {
		response.FailWithMessage("删除文件失败", c)
		global.TD27_LOG.Error("删除文件失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除文件成功", c)
	}
}
