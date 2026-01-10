package authority

import (
	entityAuthority "server/internal/model/authority/api"
	authorityResp "server/internal/model/authority/response"
	"server/internal/model/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	serviceAuthority "server/internal/service/authority"
)

type ApiApi struct {
	apiService *serviceAuthority.ApiService
}

func NewApiApi() *ApiApi {
	return &ApiApi{apiService: serviceAuthority.NewApiService()}
}

// Create
// @Tags      Create
// @Summary   添加api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      entityAuthority.ApiModel true "请求参数"
// @Success   200   {object}  commonResp.Response{data=entityAuthority.ApiModel,msg=string}
// @Router    /api/create [post]
func (aa *ApiApi) Create(c *gin.Context) {
	var apiModel entityAuthority.ApiModel
	if err := c.ShouldBindJSON(&apiModel); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := aa.apiService.Create(&apiModel); err != nil {
		common.Fail(c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
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
// @Param     data  body      authorityReq.ApiSearchParams true  "请求参数"
// @Success   200   {object}  commonResp.Response{data=commonResp.Page{list=[]entityAuthority.ApiModel},msg=string}
// @Router    /api/list [post]
func (aa *ApiApi) List(c *gin.Context) {
	var apiSp entityAuthority.ApiSearchParams
	if err := c.ShouldBindJSON(&apiSp); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := aa.apiService.List(apiSp); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			List:     list,
			Total:    total,
			Page:     apiSp.Page,
			PageSize: apiSp.PageSize,
		}, "获取成功", c)
	}
}

// Delete
// @Tags      Delete
// @Summary   删除api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CId true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /api/delete [post]
func (aa *ApiApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := aa.apiService.Delete(cId.ID); err != nil {
		common.Fail(c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
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
// @Param     data  body      commonReq.CIds true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /api/deleteApiById [post]
func (aa *ApiApi) DeleteByIds(c *gin.Context) {
	var cIds common.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := aa.apiService.DeleteByIds(cIds.IDs); err != nil {
		common.Fail(c)
		global.TD27_LOG.Error("批量删除失败", zap.Error(err))
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
// @Param     data  body      entityAuthority.ApiModel true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /api/editApi [post]
func (aa *ApiApi) Update(c *gin.Context) {
	var apiModel entityAuthority.ApiModel
	if err := c.ShouldBindJSON(&apiModel); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := aa.apiService.Update(&apiModel); err != nil {
		common.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

// GetElTree
// @Tags      GetElTree
// @Summary   格式化列出所有api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CId true "请求参数"
// @Success   200   {object}  commonResp.Response{data=authorityResp.ApiTree{list=[]entityAuthority.ApiTree,checkedKey=[]string},msg=string}
// @Router    /api/getElTreeApis [post]
func (aa *ApiApi) GetElTree(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	list, checkedKey, err := aa.apiService.GetElTree(cId.ID)
	if err != nil {
		common.FailWithMessage("获取失败", c)
	} else {
		common.OkWithDetailed(authorityResp.ApiTree{
			List:       list,
			CheckedKey: checkedKey,
		}, "获取成功", c)
	}
}
