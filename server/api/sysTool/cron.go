package sysTool

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/global"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
	modelSysTool "server/model/sysTool"
	sysToolReq "server/model/sysTool/request"
)

type CronApi struct{}

// GetCronList
// @Tags      CronApi
// @Summary   分页获取cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.PageInfo true  "page（可选）, pageSize（可选）"
// @Success   200   {object}  commonRes.Response{data=commonRes.PageResult{list=[]modelSysTool.CronModel},msg=string}
// @Router    /cron/getCronList [post]
func (st *CronApi) GetCronList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if list, total, err := cronService.GetCronList(pageInfo); err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// AddCron
// @Tags      CronApi
// @Summary   添加cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.CronModel true  "名称，方法，cron表达式，策略，开关，额外参数，备注（可选）"
// @Success   200   {object}  commonRes.Response{msg=string,data=modelSysTool.CronModel}
// @Router    /cron/addCron [post]
func (st *CronApi) AddCron(c *gin.Context) {
	var cronModel modelSysTool.CronModel
	if err := c.ShouldBindJSON(&cronModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if cron, err := cronService.AddCron(&cronModel); err != nil {
		commonRes.FailWithMessage("创建失败", c)
		global.TD27_LOG.Error("创建失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(cron, "创建成功", c)
	}
}

// DeleteCron
// @Tags      CronApi
// @Summary   删除cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CId true  "id"
// @Success   200   {object}  commonRes.Response{msg=string}
// @Router    /cron/deleteCron [post]
func (st *CronApi) DeleteCron(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := cronService.DeleteCron(cId.ID); err != nil {
		commonRes.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}

// DeleteCronByIds 批量删除cron
func (st *CronApi) DeleteCronByIds(c *gin.Context) {
	var cIds commonReq.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := cronService.DeleteCronByIds(cIds.IDs); err != nil {
		commonRes.FailWithMessage("批量删除失败", c)
		global.TD27_LOG.Error("批量删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("批量删除成功", c)
	}
}

// EditCron
// @Tags      CronApi
// @Summary   编辑cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysTool.CronModel true  "id（必须），名称，方法，cron表达式，策略，开关，额外参数，备注"
// @Success   200   {object}  commonRes.Response{msg=string,data=modelSysTool.CronModel}
// @Router    /cron/editCron [post]
func (st *CronApi) EditCron(c *gin.Context) {
	var cronModel modelSysTool.CronModel
	if err := c.ShouldBindJSON(&cronModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if instance, err := cronService.EditCron(&cronModel); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(instance, "编辑成功", c)
	}
}

// SwitchOpen
// @Tags      CronApi
// @Summary   开关cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      sysToolReq.SwitchReq true  "id, open"
// @Success   200   {object}  commonRes.Response{data=map[string]int, msg=string}
// @Router    /cron/switchOpen [post]
func (st *CronApi) SwitchOpen(c *gin.Context) {
	var switchReq sysToolReq.SwitchReq
	if err := c.ShouldBindJSON(&switchReq); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if entryId, err := cronService.SwitchOpen(switchReq.ID, switchReq.Open); err != nil {
		commonRes.FailWithMessage("切换失败", c)
		global.TD27_LOG.Error("切换失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(gin.H{"entryId": entryId}, "切换成功", c)
	}
}
