package sysTool

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysTool "server/internal/model/sysTool"
	serviceSysTool "server/internal/service/sysTool"
)

type FileApi struct {
	fileService *serviceSysTool.FileService
}

func NewFileApi() *FileApi {
	return &FileApi{
		fileService: serviceSysTool.NewFileService(),
	}
}

// Upload
// @Tags      FileApi
// @Summary   上传文件
// @Security  ApiKeyAuth
// @accept    mpfd
// @Produce   application/json
// @Param     file formData file true "The file to upload"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /file/upload [post]
func (a *FileApi) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.FailWithStatusMessage(400, fmt.Sprintf("上传失败：%s", err.Error()), c)
		return
	}

	// 只允许上传 csv文件
	if file.Header.Get("Content-Type") != "text/csv" {
		common.FailWithStatusMessage(400, "只允许上传 csv文件", c)
		return
	}

	if fullInfo, err := a.fileService.Upload(file); err != nil {
		common.FailWithStatusMessage(400, "上传失败", c)
		global.TD27_LOG.Error("上传失败", zap.Error(err))
	} else {
		common.OkWithDetailed(fullInfo, "上传成功", c)
	}
}

// List
// @Tags      FileApi
// @Summary   分页获取文件信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.ListFileReq true  "请求参数"
// @Success   200   {object}  common.Response{data=[]modelSysTool.FileModel,msg=string}
// @Router    /file/list [post]
func (a *FileApi) List(c *gin.Context) {
	var req modelSysTool.ListFileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := a.fileService.List(req); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
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
func (a *FileApi) Download(c *gin.Context) {
	fileName := c.Query("name")

	path := fmt.Sprintf("%s/%s", global.TD27_CONFIG.File.Upload, fileName)

	// 打开文件
	_, err := os.Stat(path)
	if err != nil {
		common.FailWithMessage(fmt.Sprintf("文件错误：%v", err), c)
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
// @Success   200   {object}  common.Response{msg=string}
// @Router    /file/delete [get]
func (a *FileApi) Delete(c *gin.Context) {
	fileName := c.Query("name")

	if err := a.fileService.Delete(fileName); err != nil {
		common.FailWithMessage("删除文件失败", c)
		global.TD27_LOG.Error("删除文件失败", zap.Error(err))
	} else {
		common.OkWithMessage("删除文件成功", c)
	}
}
