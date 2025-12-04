package sysTool

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	commonReq "server/internal/model/common/request"
	commonResp "server/internal/model/common/response"
	entitySysTool "server/internal/model/entity/sysTool"
	sysToolReq "server/internal/model/entity/sysTool/request"
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
	var pageInfo commonReq.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if list, total, err := ca.cronService.GetCronList(pageInfo); err != nil {
		commonResp.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(commonResp.Page{
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
		commonResp.FailReq(err.Error(), c)
		return
	}

	if cron, err := ca.cronService.AddCron(&cronModel); err != nil {
		commonResp.FailWithMessage("创建失败", c)
		global.TD27_LOG.Error("创建失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(cron, "创建成功", c)
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
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ca.cronService.DeleteCron(cId.ID); err != nil {
		commonResp.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonResp.OkWithMessage("删除成功", c)
	}
}

// DeleteCronByIds 批量删除cron
func (ca *CronApi) DeleteCronByIds(c *gin.Context) {
	var cIds commonReq.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ca.cronService.DeleteCronByIds(cIds.IDs); err != nil {
		commonResp.FailWithMessage("批量删除失败", c)
		global.TD27_LOG.Error("批量删除失败", zap.Error(err))
	} else {
		commonResp.OkWithMessage("批量删除成功", c)
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
		commonResp.FailReq(err.Error(), c)
		return
	}

	if instance, err := ca.cronService.EditCron(&cronModel); err != nil {
		commonResp.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(instance, "编辑成功", c)
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
		commonResp.FailReq(err.Error(), c)
		return
	}

	if entryId, err := ca.cronService.SwitchOpen(switchReq.ID, switchReq.Open); err != nil {
		commonResp.FailWithMessage("切换失败", c)
		global.TD27_LOG.Error("切换失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(gin.H{"entryId": entryId}, "切换成功", c)
	}
}
