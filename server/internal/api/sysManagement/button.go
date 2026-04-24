package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	"server/internal/service/sysManagement"
)

type ButtonApi struct {
	service *sysManagement.ButtonService
}

func NewButtonApi() *ButtonApi {
	return &ButtonApi{service: sysManagement.NewButtonService()}
}

func (a *ButtonApi) Create(c *gin.Context) {
	var req modelSysManagement.CreateButtonReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}
	button, err := a.service.Create(&req)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("创建按钮失败", "error", err)
		return
	}
	common.OkWithData(button, c)
}

func (a *ButtonApi) Update(c *gin.Context) {
	var req modelSysManagement.UpdateButtonReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}
	if err := a.service.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("更新按钮失败", "error", err)
		return
	}
	common.Ok(c)
}

func (a *ButtonApi) Delete(c *gin.Context) {
	var req common.CId
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}
	if err := a.service.Delete(req.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}

func (a *ButtonApi) List(c *gin.Context) {
	var req modelSysManagement.ListButtonReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}
	list, total, err := a.service.List(&req)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.OkWithData(map[string]interface{}{"list": list, "total": total}, c)
}

func (a *ButtonApi) GetPageButtons(c *gin.Context) {
	pagePath := c.Query("pagePath")
	if pagePath == "" {
		common.FailReq("pagePath is required", c)
		return
	}
	claims, err := GetClaims(c)
	if err != nil {
		common.FailWithDetailed(gin.H{"reload": true}, "未登录", c)
		c.Abort()
		return
	}
	buttons, err := a.service.GetPageButtons(pagePath, claims.GetAllRoleIDs())
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.OkWithData(buttons, c)
}

func (a *ButtonApi) CheckPermission(c *gin.Context) {
	var req modelSysManagement.CheckButtonReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}
	claims, err := GetClaims(c)
	if err != nil {
		common.FailWithDetailed(gin.H{"reload": true}, "未登录", c)
		c.Abort()
		return
	}
	hasPermission := a.service.CheckPermission(req.ButtonCode, claims.GetAllRoleIDs())
	common.OkWithData(hasPermission, c)
}

func (a *ButtonApi) BatchCheckPermission(c *gin.Context) {
	var req struct {
		ButtonCodes []string `json:"buttonCodes" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}
	claims, err := GetClaims(c)
	if err != nil {
		common.FailWithDetailed(gin.H{"reload": true}, "未登录", c)
		c.Abort()
		return
	}
	result := a.service.BatchCheckPermission(req.ButtonCodes, claims.GetAllRoleIDs())
	common.OkWithData(result, c)
}

func (a *ButtonApi) GetUserButtons(c *gin.Context) {
	claims, err := GetClaims(c)
	if err != nil {
		common.FailWithDetailed(gin.H{"reload": true}, "未登录", c)
		c.Abort()
		return
	}
	buttons, err := a.service.GetUserButtons(claims.GetAllRoleIDs())
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.OkWithData(buttons, c)
}
