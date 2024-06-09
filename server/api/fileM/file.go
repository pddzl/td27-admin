package fileM

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"

	"server/global"
	commonRes "server/model/common/response"
	fileMReq "server/model/fileM/request"
)

type FileApi struct{}

// Upload
// @Tags      FileApi
// @Summary   上传文件
// @Security  ApiKeyAuth
// @accept    mpfd
// @Produce   application/json
// @Param     file formData file true "The file to upload"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /file/upload [post]
func (f *FileApi) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		commonRes.FailWithStatusMessage(400, fmt.Sprintf("上传失败：%s", err.Error()), c)
		return
	}

	// 只允许上传csv文件
	if file.Header.Get("Content-Type") != "text/csv" {
		commonRes.FailWithStatusMessage(400, "只允许上传csv文件", c)
		return
	}

	if fullInfo, err := fileService.Upload(file); err != nil {
		commonRes.FailWithStatusMessage(400, "上传失败", c)
		global.TD27_LOG.Error("上传失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(fullInfo, "上传成功", c)
	}
}

// GetFileList
// @Tags      FileApi
// @Summary   分页获取文件信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      fileMReq.FileSearchParams true  "请求参数"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]fileM.FileModel},msg=string}
// @Router    /file/getFileList [post]
func (f *FileApi) GetFileList(c *gin.Context) {
	var params fileMReq.FileSearchParams
	if err := c.ShouldBindJSON(&params); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := fileService.GetFileList(params); err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
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
func (f *FileApi) Download(c *gin.Context) {
	fileName := c.Query("name")

	path := fmt.Sprintf("%s/%s", global.TD27_CONFIG.File.Upload, fileName)

	// 打开文件
	_, err := os.Stat(path)
	if err != nil {
		commonRes.FailWithMessage(fmt.Sprintf("文件错误：%v", err), c)
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
func (f *FileApi) Delete(c *gin.Context) {
	fileName := c.Query("name")

	if err := fileService.Delete(fileName); err != nil {
		commonRes.FailWithMessage("删除文件失败", c)
		global.TD27_LOG.Error("删除文件失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除文件成功", c)
	}
}
