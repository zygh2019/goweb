package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code, "data": data, "msg": msg,
	})
}

const (
	success = 0
	error   = 500
)
const (
	okMsg = "成功"
)

func Ok(c *gin.Context) {
	//空对象
	Result(success, map[string]any{}, ErrorMap[ErrorCode(success)], c)
}
func failWithMsg(msg string, c *gin.Context) {
	//空对象
	Result(error, map[string]any{}, msg, c)
}

func failWithCode(code int, msg string, c *gin.Context) {
	//空对象
	Result(code, map[string]any{}, ErrorMap[ErrorCode(code)], c)
}
func OkWithData(data any, c *gin.Context) {
	Result(success, data, ErrorMap[ErrorCode(success)], c)
}
