package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
	"log/slog"
)

type ApiApi struct {
	apiService *serviceSysManagement.ApiService
}

func NewApiApi() *ApiApi {
	return &ApiApi{apiService: serviceSysManagement.NewApiService()}
}

// Create
// @Tags      Create
// @Summary   添加api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.CreateApiReq true "请求参数"
// @Success   200   {object}  common.Response{data=modelSysManagement.ApiModel,msg=string}
// @Router    /api/create [post]
func (a *ApiApi) Create(c *gin.Context) {
	var req modelSysManagement.CreateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := a.apiService.Create(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("添加失败", "error", err)
	} else {
		common.OkWithDetailed(instance, "添加成功", c)
	}
}

// List
// @Tags      List
// @Summary   分页获取api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.ListApiReq true  "请求参数"
// @Success   200   {object}  common.Response{data=common.Page{list=[]modelSysManagement.ApiModel},msg=string}
// @Router    /api/list [post]
func (a *ApiApi) List(c *gin.Context) {
	var req modelSysManagement.ListApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := a.apiService.List(&req); err != nil {
		common.FailWithMessage("获取失败", c)
		slog.Error("获取失败", "error", err)
	} else {
		common.OkWithDetailed(common.Page{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// Delete
// @Tags      Delete
// @Summary   删除api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /api/delete [post]
func (a *ApiApi) Delete(c *gin.Context) {
	var req common.CId
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.apiService.Delete(req.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("删除失败", "error", err)
	} else {
		common.Ok(c)
	}
}

// DeleteByIds
// @Tags      DeleteByIds
// @Summary   批量删除api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CIds true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /api/deleteByIds [post]
func (a *ApiApi) DeleteByIds(c *gin.Context) {
	var req common.CIds
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.apiService.DeleteByIds(req.IDs); err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("批量删除失败", "error", err)
	} else {
		common.Ok(c)
	}
}

// Update
// @Tags      Update
// @Summary   编辑api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.UpdateApiReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /api/update [post]
func (a *ApiApi) Update(c *gin.Context) {
	var req modelSysManagement.UpdateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	_, err := a.apiService.Update(&req)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("更新失败", "error", err)
	} else {
		common.Ok(c)
	}
}

// ElTree
// @Tags      ElTree
// @Summary   格式化列出所有api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.ApiTreeReq true "请求参数"
// @Success   200   {object}  common.Response{data=modelSysManagement.ApiTreeResp{list=[]modelSysManagement.ApiTreeNode,checkedKey=[]string,checkedIds=[]uint},msg=string}
// @Router    /api/elTree [post]
func (a *ApiApi) ElTree(c *gin.Context) {
	var req modelSysManagement.ApiTreeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	list, checkedIds, err := a.apiService.ElTree(&req)
	if err != nil {
		common.FailWithMessage("获取失败", c)
	} else {
		common.OkWithDetailed(modelSysManagement.ApiTreeResp{
			List:       list,
			CheckedIds: checkedIds,
		}, "获取成功", c)
	}
}
