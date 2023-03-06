package middleware

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/common/response"
	"server/service"
	"server/utils"
	"strconv"
)

var (
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.TD27_CONFIG.System.Env != "develop" {
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
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
