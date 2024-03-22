package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/global"
	modelAuthority "server/model/authority"
	authorityReq "server/model/authority/request"
	authorityRes "server/model/authority/response"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
)

type ApiApi struct{}

// AddApi
// @Tags      ApiApi
// @Summary   添加api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelAuthority.ApiModel true "请求参数"
// @Success   200   {object}  response.Response{data=modelAuthority.ApiModel,msg=string}
// @Router    /api/addApi [post]
func (a *ApiApi) AddApi(c *gin.Context) {
	var apiModel modelAuthority.ApiModel
	if err := c.ShouldBindJSON(&apiModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if instance, err := apiService.AddApi(&apiModel); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(instance, "添加成功", c)
	}
}

// GetApis
// @Tags      ApiApi
// @Summary   分页获取api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.ApiSearchParams true  "请求参数"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]modelAuthority.ApiModel},msg=string}
// @Router    /api/getApis [post]
func (a *ApiApi) GetApis(c *gin.Context) {
	var apiSp authorityReq.ApiSearchParams
	if err := c.ShouldBindJSON(&apiSp); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if list, total, err := apiService.GetApis(apiSp); err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
			List:     list,
			Total:    total,
			Page:     apiSp.Page,
			PageSize: apiSp.PageSize,
		}, "获取成功", c)
	}
}

// DeleteApi
// @Tags      ApiApi
// @Summary   删除api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/deleteApi [post]
func (a *ApiApi) DeleteApi(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := apiService.DeleteApi(cId.ID); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}

// DeleteApiById
// @Tags      ApiApi
// @Summary   批量删除api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CIds true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/deleteApiById [post]
func (a *ApiApi) DeleteApiById(c *gin.Context) {
	var cIds commonReq.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := apiService.DeleteApiById(cIds.IDs); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("批量删除失败", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}

// EditApi
// @Tags      ApiApi
// @Summary   编辑api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelAuthority.ApiModel true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/editApi [post]
func (a *ApiApi) EditApi(c *gin.Context) {
	var apiModel modelAuthority.ApiModel
	if err := c.ShouldBindJSON(&apiModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := apiService.EditApi(&apiModel); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}

// GetElTreeApis
// @Tags      ApiApi
// @Summary   格式化列出所有api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{data=authorityRes.ApiTree{list=[]modelAuthority.ApiTree,checkedKey=[]string},msg=string}
// @Router    /api/getElTreeApis [post]
func (a *ApiApi) GetElTreeApis(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	list, checkedKey, err := apiService.GetElTreeApis(cId.ID)
	if err != nil {
		commonRes.FailWithMessage("获取失败", c)
	} else {
		commonRes.OkWithDetailed(authorityRes.ApiTree{
			List:       list,
			CheckedKey: checkedKey,
		}, "获取成功", c)
	}
}
