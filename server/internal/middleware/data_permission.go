package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	apiSysManagement "server/internal/api/sysManagement"
	"server/internal/global"
)

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
