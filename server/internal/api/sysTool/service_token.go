package sysTool

import (
	"github.com/gin-gonic/gin"

	"server/internal/model/common"
	modelSysTool "server/internal/model/sysTool"
	"server/internal/service/sysTool"
	"log/slog"
)

type ServiceTokenApi struct {
	service *sysTool.ServiceTokenService
}

func NewServiceTokenApi() *ServiceTokenApi {
	return &ServiceTokenApi{
		service: sysTool.NewServiceTokenService(),
	}
}

// Create
// @Tags      ServiceToken
// @Summary   创建服务令牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.CreateServiceTokenReq  true  "创建参数"
// @Success   200   {object}  common.Response{data=modelSysTool.CreateServiceTokenResp}
// @Router    /serviceToken/create [post]
func (a *ServiceTokenApi) Create(c *gin.Context) {
	var req modelSysTool.CreateServiceTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	slog.Debug("API Create被调用",
		"name", req.Name,
		"apiIDs", req.ApiIDs)

	resp, err := a.service.Create(&req)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("创建服务令牌失败", "error", err)
		return
	}

	common.OkWithData(resp, c)
}

// Update
// @Tags      ServiceToken
// @Summary   更新服务令牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.UpdateServiceTokenReq  true  "更新参数"
// @Success   200   {object}  common.Response
// @Router    /serviceToken/update [post]
func (a *ServiceTokenApi) Update(c *gin.Context) {
	var req modelSysTool.UpdateServiceTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.service.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("更新服务令牌失败", "error", err)
		return
	}

	common.Ok(c)
}

// Delete
// @Tags      ServiceToken
// @Summary   删除服务令牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId  true  "ID"
// @Success   200   {object}  common.Response
// @Router    /serviceToken/delete [post]
func (a *ServiceTokenApi) Delete(c *gin.Context) {
	var req common.CId
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.service.Delete(req.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("删除服务令牌失败", "error", err)
		return
	}

	common.Ok(c)
}

// GetById
// @Tags      ServiceToken
// @Summary   获取服务令牌详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     common.CId  true  "ID"
// @Success   200   {object}  common.Response{data=modelSysTool.ServiceTokenDetailResp}
// @Router    /serviceToken/detail [post]
func (a *ServiceTokenApi) GetById(c *gin.Context) {
	var req common.CId
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	resp, err := a.service.GetByID(req.ID)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	common.OkWithData(resp, c)
}

// List
// @Tags      ServiceToken
// @Summary   获取服务令牌列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     modelSysTool.ListServiceTokenReq  true  "查询参数"
// @Success   200   {object}  common.Response{data=modelSysTool.ServiceTokenListResp}
// @Router    /serviceToken/list [post]
func (a *ServiceTokenApi) List(c *gin.Context) {
	var req modelSysTool.ListServiceTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	resp, err := a.service.List(&req)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("获取服务令牌列表失败", "error", err)
		return
	}

	common.OkWithData(resp, c)
}
