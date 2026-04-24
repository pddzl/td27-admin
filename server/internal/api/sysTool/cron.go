package sysTool

import (
	"github.com/gin-gonic/gin"

	"server/internal/model/common"
	modelSysTool "server/internal/model/sysTool"
	serviceSysTool "server/internal/service/sysTool"
	"log/slog"
)

type CronApi struct {
	cronService *serviceSysTool.CronService
}

func NewCronApi() *CronApi {
	return &CronApi{
		cronService: serviceSysTool.NewCronService(),
	}
}

// List
// @Tags      CronApi
// @Summary   分页获取cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.PageInfo true  "page（可选）, pageSize（可选）"
// @Success   200   {object}  common.Response{data=common.Page{list=[]modelSysTool.CronModel},msg=string}
// @Router    /cron/list [post]
func (a *CronApi) List(c *gin.Context) {
	var req common.PageInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := a.cronService.List(&req); err != nil {
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

// Create
// @Tags      CronApi
// @Summary   添加cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.CronModel true  "名称，方法，cron表达式，策略，开关，额外参数，备注（可选）"
// @Success   200   {object}  common.Response{msg=string,data=modelSysTool.CronModel}
// @Router    /cron/create [post]
func (a *CronApi) Create(c *gin.Context) {
	var req modelSysTool.CronModel
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if cron, err := a.cronService.Create(&req); err != nil {
		common.FailWithMessage("创建失败", c)
		slog.Error("创建失败", "error", err)
	} else {
		common.OkWithDetailed(cron, "创建成功", c)
	}
}

// Delete
// @Tags      CronApi
// @Summary   删除cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true  "id"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /cron/delete [post]
func (a *CronApi) Delete(c *gin.Context) {
	var req common.CId
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.cronService.Delete(req.ID); err != nil {
		common.FailWithMessage("删除失败", c)
		slog.Error("删除失败", "error", err)
	} else {
		common.OkWithMessage("删除成功", c)
	}
}

func (a *CronApi) DeleteByIds(c *gin.Context) {
	var req common.CIds
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.cronService.DeleteByIds(req.IDs); err != nil {
		common.FailWithMessage("批量删除失败", c)
		slog.Error("批量删除失败", "error", err)
	} else {
		common.OkWithMessage("批量删除成功", c)
	}
}

// Update
// @Tags      CronApi
// @Summary   编辑cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.CronModel true  "id（必须），名称，方法，cron表达式，策略，开关，额外参数，备注"
// @Success   200   {object}  common.Response{msg=string,data=modelSysTool.CronModel}
// @Router    /cron/update [post]
func (a *CronApi) Update(c *gin.Context) {
	var req modelSysTool.CronModel
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := a.cronService.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		slog.Error("编辑失败", "error", err)
	} else {
		common.OkWithDetailed(instance, "编辑成功", c)
	}
}

// SwitchOpen
// @Tags      CronApi
// @Summary   开关cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.SwitchReq true  "id, open"
// @Success   200   {object}  common.Response{data=map[string]int, msg=string}
// @Router    /cron/switchOpen [post]
func (a *CronApi) SwitchOpen(c *gin.Context) {
	var req modelSysTool.SwitchReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if entryId, err := a.cronService.SwitchOpen(req.ID, req.Open); err != nil {
		common.FailWithMessage("切换失败", c)
		slog.Error("切换失败", "error", err)
	} else {
		common.OkWithDetailed(gin.H{"entryId": entryId}, "切换成功", c)
	}
}

// RunOnce 立即执行一次
// @Tags      CronApi
// @Summary   立即执行定时任务
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true  "id"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /cron/runOnce [post]
func (a *CronApi) RunOnce(c *gin.Context) {
	var req common.CId
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.cronService.RunOnce(req.ID); err != nil {
		common.FailWithMessage("执行失败", c)
		slog.Error("执行失败", "error", err)
	} else {
		common.OkWithMessage("任务已触发", c)
	}
}
