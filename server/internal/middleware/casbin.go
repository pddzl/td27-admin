package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"

	apiSysManagement "server/internal/api/sysManagement"
	"server/internal/global"
	"server/internal/model/common"
	"server/internal/service/sysManagement"
)

var (
	casbinService = sysManagement.NewCasbinService()
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.TD27_CONFIG.System.Env != "dev" {
			waitUse, _ := apiSysManagement.GetClaims(c)
			//获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			// 角色ID
			sub := strconv.Itoa(int(waitUse.RoleId))
			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				common.FailWithDetailed(gin.H{}, "接口权限不足", c)
				global.TD27_LOG.Error("接口权限不足")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
