package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/utils"
)

type UserApi struct{}

func (u *UserApi) GetUserInfo(c *gin.Context) {
	if userInfo, err := utils.GetUserInfo(c); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		response.OkWithDetailed(userInfo, "获取成功", c)
	}
}
