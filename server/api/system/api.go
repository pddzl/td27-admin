package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
	systemRep "server/model/system/response"
)

type ApiApi struct{}

// AddApi
// @Tags      ApiApi
// @Summary   添加api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemModel.ApiModel true "请求参数"
// @Success   200   {object}  response.Response{data=systemModel.ApiModel,msg=string}
// @Router    /api/addApi [post]
func (a *ApiApi) AddApi(c *gin.Context) {
	var apiReq systemModel.ApiModel
	_ = c.ShouldBindJSON(&apiReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&apiReq); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if instance, err := apiService.AddApi(apiReq); err != nil {
		response.FailWithMessage("添加失败", c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		response.OkWithDetailed(instance, "添加成功", c)
	}
}

// GetApis
// @Tags      ApiApi
// @Summary   分页获取api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.ApiSearchParams true  "请求参数"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]system.ApiModel},msg=string}
// @Router    /api/getApis [post]
func (a *ApiApi) GetApis(c *gin.Context) {
	var apiSp systemReq.ApiSearchParams
	_ = c.ShouldBindJSON(&apiSp)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&apiSp); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if list, total, err := apiService.GetApis(apiSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
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
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.DeleteApi(cId.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
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
	var cIds request.CIds
	_ = c.ShouldBindJSON(&cIds)

	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cIds); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.DeleteApiById(cIds.Ids); err != nil {
		response.FailWithMessage("批量删除失败", c)
		global.TD27_LOG.Error("批量删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// EditApi
// @Tags      ApiApi
// @Summary   编辑api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.EditApi true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/editApi [post]
func (a *ApiApi) EditApi(c *gin.Context) {
	var eApi systemReq.EditApi
	_ = c.ShouldBindJSON(&eApi)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&eApi); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.EditApi(eApi); err != nil {
		response.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

// GetElTreeApis
// @Tags      ApiApi
// @Summary   格式化列出所有api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{data=systemRep.ApiTree{list=[]systemModel.ApiTree,checkedKey=[]string},msg=string}
// @Router    /api/getElTreeApis [post]
func (a *ApiApi) GetElTreeApis(c *gin.Context) {
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	list, checkedKey, err := apiService.GetElTreeApis(cId.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRep.ApiTree{
			List:       list,
			CheckedKey: checkedKey,
		}, "获取成功", c)
	}
}
