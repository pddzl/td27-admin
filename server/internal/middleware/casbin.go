package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	apiSysManagement "server/internal/api/sysManagement"
	"server/internal/global"
	"server/internal/model/common"
	serviceSysManagement "server/internal/service/sysManagement"
)

var (
	casbinService         = serviceSysManagement.NewCasbinService()
	dataPermissionService = serviceSysManagement.NewDataPermissionService()
)

// CasbinHandler 拦截器（支持多角色和权限缓存）
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开发环境跳过权限检查
		if global.TD27_CONFIG.System.Env == "dev" {
			c.Next()
			return
		}

		claims, err := apiSysManagement.GetClaims(c)
		if err != nil {
			common.FailWithDetailed(gin.H{}, "未登录或token无效", c)
			c.Abort()
			return
		}

		// 获取请求的PATH和方法
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 获取用户的所有角色ID
		roleIDs := claims.GetAllRoleIDs()

		// checking by Casbin
		success, err := casbinService.Enforce(roleIDs, obj, act)
		if err != nil {
			global.TD27_LOG.Error("权限检查失败",
				zap.Error(err),
				zap.Uints("roleIDs", roleIDs),
				zap.String("path", obj),
				zap.String("method", act))
			common.FailWithDetailed(gin.H{}, "权限检查失败", c)
			c.Abort()
			return
		}

		if !success {
			common.FailWithDetailed(gin.H{}, "接口权限不足", c)
			global.TD27_LOG.Warn("接口权限不足",
				zap.Uint("userID", claims.ID),
				zap.Uints("roleIDs", roleIDs),
				zap.String("path", obj),
				zap.String("method", act))
			c.Abort()
			return
		}

		c.Next()
	}
}

// DataPermissionHandler 数据权限中间件
func DataPermissionHandler(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !global.TD27_CONFIG.Casbin.EnableDataPermission {
			c.Next()
			return
		}

		claims, err := apiSysManagement.GetClaims(c)
		if err != nil {
			c.Next()
			return
		}

		// 将数据权限信息存入上下文，供后续使用
		dataPerm, err := dataPermissionService.GetUserDataPermission(c, claims.ID, resource)
		if err != nil {
			global.TD27_LOG.Error("获取数据权限失败", zap.Error(err))
		}

		c.Set("dataPermission", dataPerm)
		c.Next()
	}
}
