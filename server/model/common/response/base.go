package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR_RES = 7 // 响应错误
	ERROR_REQ = 4 // 请求错误
	SUCCESS   = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ResultStatus(status int, code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(status, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR_RES, map[string]interface{}{}, "操作失败", c)
}

func FailReq(message string, c *gin.Context) {
	Result(ERROR_REQ, map[string]interface{}{}, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR_RES, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR_RES, data, message, c)
}

func FailWithStatusMessage(status int, message string, c *gin.Context) {
	ResultStatus(status, ERROR_RES, map[string]interface{}{}, message, c)
}
