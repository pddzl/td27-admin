package authority

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	var apiReq modelAuthority.ApiModel
	_ = c.ShouldBindJSON(&apiReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&apiReq); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if instance, err := apiService.AddApi(apiReq); err != nil {
		commonRes.FailWithMessage("添加失败", c)
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
	_ = c.ShouldBindJSON(&apiSp)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&apiSp); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
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
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.DeleteApi(cId.ID); err != nil {
		commonRes.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除成功", c)
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
	_ = c.ShouldBindJSON(&cIds)

	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&cIds); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.DeleteApiById(cIds.Ids); err != nil {
		commonRes.FailWithMessage("批量删除失败", c)
		global.TD27_LOG.Error("批量删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("批量删除成功", c)
	}
}

// EditApi
// @Tags      ApiApi
// @Summary   编辑api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditApi true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/editApi [post]
func (a *ApiApi) EditApi(c *gin.Context) {
	var eApi authorityReq.EditApi
	_ = c.ShouldBindJSON(&eApi)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&eApi); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.EditApi(eApi); err != nil {
		commonRes.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("编辑成功", c)
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
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
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
