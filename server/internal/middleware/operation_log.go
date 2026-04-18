package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	modelSysMonitor "server/internal/model/sysMonitor"
	"server/internal/pkg/async"
)

// GetAsyncLoggerBufferLen 获取异步日志缓冲区长度（用于监控）
func GetAsyncLoggerBufferLen() int {
	return asyncLogger.BufferLen()
}

var (
	asyncLogger = async.GetAsyncLogger()
)

type responseProxyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseProxyWriter) Write(b []byte) (int, error) {
	if r.body == nil {
		r.body = bytes.NewBufferString("")
	}
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求参数
		var reqParam []byte
		if c.Request.Method == http.MethodGet {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			reqParam, _ = json.Marshal(&m)
		} else {
			var err error
			reqParam, err = io.ReadAll(c.Request.Body)
			if err != nil {
				// 读取请求体失败，继续处理
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(reqParam))
			}
		}

		// 解析token
		claims, err := sysManagement.GetClaims(c)
		
		record := modelSysMonitor.OperationLogModel{
			Ip:        c.ClientIP(),
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			UserAgent: c.Request.UserAgent(),
			ReqParam:  string(reqParam),
		}
		
		if err == nil && claims != nil {
			record.UserID = claims.ID
			record.UserName = claims.Username
		} else if isServiceToken := c.GetBool("isServiceToken"); isServiceToken {
			record.UserName = "service_token"
			if prefix, exists := c.Get("serviceTokenPrefix"); exists {
				record.UserName = fmt.Sprintf("service_token(%s)", prefix)
			}
		}

		writer := responseProxyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		record.Status = c.Writer.Status()
		record.RespTime = time.Since(now).Milliseconds()
		record.RespData = writer.body.String()

		// 使用异步日志记录，不阻塞响应
		asyncLogger.Log(&record)
	}
}
