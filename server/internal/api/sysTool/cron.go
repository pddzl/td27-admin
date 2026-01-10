package sysTool

import (
	"server/internal/model/common"
	entitySysTool "server/internal/model/sysTool"
	sysToolReq "server/internal/model/sysTool/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	serviceSysTool "server/internal/service/sysTool"
)

type CronApi struct {
	cronService *serviceSysTool.CronService
}

func NewCronApi() *CronApi {
	return &CronApi{
		cronService: serviceSysTool.NewCronService(),
	}
}

// GetCronList
// @Tags      CronApi
// @Summary   分页获取cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.PageInfo true  "page（可选）, pageSize（可选）"
// @Success   200   {object}  commonResp.Response{data=commonResp.Page{list=[]entitySysTool.CronModel},msg=string}
// @Router    /cron/getCronList [post]
func (ca *CronApi) GetCronList(c *gin.Context) {
	var pageInfo common.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := ca.cronService.GetCronList(pageInfo); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
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
// @Param     data  body      entitySysTool.CronModel true  "名称，方法，cron表达式，策略，开关，额外参数，备注（可选）"
// @Success   200   {object}  commonResp.Response{msg=string,data=entitySysTool.CronModel}
// @Router    /cron/addCron [post]
func (ca *CronApi) AddCron(c *gin.Context) {
	var cronModel entitySysTool.CronModel
	if err := c.ShouldBindJSON(&cronModel); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if cron, err := ca.cronService.AddCron(&cronModel); err != nil {
		common.FailWithMessage("创建失败", c)
		global.TD27_LOG.Error("创建失败", zap.Error(err))
	} else {
		common.OkWithDetailed(cron, "创建成功", c)
	}
}

// DeleteCron
// @Tags      CronApi
// @Summary   删除cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CId true  "id"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /cron/deleteCron [post]
func (ca *CronApi) DeleteCron(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ca.cronService.DeleteCron(cId.ID); err != nil {
		common.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		common.OkWithMessage("删除成功", c)
	}
}

// DeleteCronByIds 批量删除cron
func (ca *CronApi) DeleteCronByIds(c *gin.Context) {
	var cIds common.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ca.cronService.DeleteCronByIds(cIds.IDs); err != nil {
		common.FailWithMessage("批量删除失败", c)
		global.TD27_LOG.Error("批量删除失败", zap.Error(err))
	} else {
		common.OkWithMessage("批量删除成功", c)
	}
}

// EditCron
// @Tags      CronApi
// @Summary   编辑cron
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      entitySysTool.CronModel true  "id（必须），名称，方法，cron表达式，策略，开关，额外参数，备注"
// @Success   200   {object}  commonResp.Response{msg=string,data=entitySysTool.CronModel}
// @Router    /cron/editCron [post]
func (ca *CronApi) EditCron(c *gin.Context) {
	var cronModel entitySysTool.CronModel
	if err := c.ShouldBindJSON(&cronModel); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := ca.cronService.EditCron(&cronModel); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
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
// @Param     data  body      sysToolReq.SwitchReq true  "id, open"
// @Success   200   {object}  commonResp.Response{data=map[string]int, msg=string}
// @Router    /cron/switchOpen [post]
func (ca *CronApi) SwitchOpen(c *gin.Context) {
	var switchReq sysToolReq.SwitchReq
	if err := c.ShouldBindJSON(&switchReq); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if entryId, err := ca.cronService.SwitchOpen(switchReq.ID, switchReq.Open); err != nil {
		common.FailWithMessage("切换失败", c)
		global.TD27_LOG.Error("切换失败", zap.Error(err))
	} else {
		common.OkWithDetailed(gin.H{"entryId": entryId}, "切换成功", c)
	}
}
