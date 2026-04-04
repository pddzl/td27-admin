package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type DeptApi struct {
	deptService *serviceSysManagement.DeptService
}

func NewDeptApi() *DeptApi {
	return &DeptApi{
		deptService: serviceSysManagement.NewDeptService(),
	}
}

// List
// @Tags      DeptApi
// @Summary   获取部门列表（树形）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     modelSysManagement.DeptListReq  true  "请求参数"
// @Success   200   {object}  common.Response{data=[]modelSysManagement.DeptResp}
// @Router    /dept/list [post]
func (d *DeptApi) List(c *gin.Context) {
	var req modelSysManagement.DeptListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	list, err := d.deptService.List(&req)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	common.OkWithDetailed(list, "获取成功", c)
}

// Create
// @Tags      DeptApi
// @Summary   创建部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.CreateDeptReq  true  "请求参数"
// @Success   200   {object}  common.Response
// @Router    /dept/create [post]
func (d *DeptApi) Create(c *gin.Context) {
	var req modelSysManagement.CreateDeptReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := d.deptService.Create(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	common.OkWithMessage("创建成功", c)
}

// Update
// @Tags      DeptApi
// @Summary   更新部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.UpdateDeptReq  true  "请求参数"
// @Success   200   {object}  common.Response
// @Router    /dept/update [post]
func (d *DeptApi) Update(c *gin.Context) {
	var req modelSysManagement.UpdateDeptReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := d.deptService.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	common.OkWithMessage("更新成功", c)
}

// Delete
// @Tags      DeptApi
// @Summary   删除部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId  true  "请求参数"
// @Success   200   {object}  common.Response
// @Router    /dept/delete [post]
func (d *DeptApi) Delete(c *gin.Context) {
	var req common.CId
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := d.deptService.Delete(req.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	common.OkWithMessage("删除成功", c)
}

// GetElTreeDepts
// @Tags      DeptApi
// @Summary   获取部门树（用于选择器）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  common.Response{data=modelSysManagement.DeptTreeResp}
// @Router    /dept/getElTreeDepts [post]
func (d *DeptApi) GetElTreeDepts(c *gin.Context) {
	tree, ids, err := d.deptService.GetElTreeDepts()
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	common.OkWithDetailed(gin.H{
		"tree": tree,
		"ids":  ids,
	}, "获取成功", c)
}
