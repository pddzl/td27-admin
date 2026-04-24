package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	apiSysManagement "server/internal/api/sysManagement"

	"server/internal/global"
	"server/internal/model/common"
	serviceSysManagement "server/internal/service/sysManagement"
	"log/slog"
)

var (
	casbinService         = serviceSysManagement.NewCasbinService()
	dataPermissionService = serviceSysManagement.NewDataPermissionService()
)

// CasbinHandler 拦截器（支持JWT用户角色和服务令牌）
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开发环境跳过权限检查
		if global.TD27_CONFIG.System.Env == "dev" {
			c.Next()
			return
		}

		obj := c.Request.URL.Path
		act := c.Request.Method

		if isServiceToken := c.GetBool("isServiceToken"); isServiceToken {
			// 服务令牌授权
			if tokenID, exists := c.Get("serviceTokenID"); exists {
				subject := fmt.Sprintf("token:%d", tokenID)
				success, err := casbinService.EnforceSubject(subject, obj, act)
				if err != nil {
					slog.Error("服务令牌权限检查失败",
						"error", err,
						"subject", subject,
						"path", obj,
						"method", act)
					common.FailWithDetailed(gin.H{}, "权限检查失败", c)
					c.Abort()
					return
				}

				if !success {
					common.FailWithDetailed(gin.H{}, "接口权限不足", c)
					slog.Warn("服务令牌权限不足",
						"subject", subject,
						"path", obj,
						"method", act)
					c.Abort()
					return
				}

				c.Next()
				return
			}
		} else {
			// JWT认证（用户角色权限）
			claims, err := apiSysManagement.GetClaims(c)
			if err == nil {
				roleIDs := claims.GetAllRoleIDs()
				success, err := casbinService.Enforce(roleIDs, obj, act)
				if err != nil {
					slog.Error("权限检查失败",
						"error", err,
						"roleIDs", roleIDs,
						"path", obj,
						"method", act)
					common.FailWithDetailed(gin.H{}, "权限检查失败", c)
					c.Abort()
					return
				}

				if !success {
					common.FailWithDetailed(gin.H{}, "接口权限不足", c)
					slog.Warn("接口权限不足",
						"userID", claims.ID,
						"roleIDs", roleIDs,
						"path", obj,
						"method", act)
					c.Abort()
					return
				}

				c.Next()
				return
			}
		}

		// 失败，返回未登录
		common.FailWithDetailed(gin.H{}, "未登录或token无效", c)
		c.Abort()
	}
}
