package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	systemReq "server/model/system/request"
	"server/utils"
)

type RoleApi struct{}

func (ra *RoleApi) GetRoles(c *gin.Context) {
	userInfo, err := utils.GetUserInfo(c)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	if list, err := roleService.GetRoles(userInfo.Username); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

func (ra *RoleApi) AddRole(c *gin.Context) {
	var roleReq systemReq.Role
	_ = c.ShouldBindJSON(&roleReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&roleReq); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	userInfo, err := utils.GetUserInfo(c)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	if role, err := roleService.AddRole(userInfo.Username, roleReq.RoleName); err != nil {
		response.FailWithMessage("添加失败", c)
		global.TD27_LOG.Error("添加角色失败", zap.Error(err))
	} else {
		response.OkWithDetailed(role, "添加成功", c)
	}
}
