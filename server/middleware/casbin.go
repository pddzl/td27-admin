package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"server/global"
	commonRes "server/model/common/response"
	"server/service"
	"server/utils"
)

var (
	casbinService = service.ServiceGroupApp.Base.CasbinService
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.TD27_CONFIG.System.Env != "dev" {
			waitUse, _ := utils.GetClaims(c)
			//获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			// 角色ID
			sub := strconv.Itoa(int(waitUse.RoleId))
			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				commonRes.FailWithDetailed(gin.H{}, "接口权限不足", c)
				global.TD27_LOG.Error("接口权限不足")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
