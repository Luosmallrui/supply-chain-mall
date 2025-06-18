package ctx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
	Meta    any         `json:"metadata,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {

	resp := Response{
		Code:    code,
		Message: message,
		Data:    nil,
		Meta:    initMeta(),
	}
	if resp.Data == nil {
		resp.Data = ""
	}

	c.JSON(code, resp)
}

func initMeta() map[string]any {
	meta := make(map[string]any)
	pc, _, line, ok := runtime.Caller(2)
	if ok {
		meta["error_line"] = line
		meta["error_function"] = runtime.FuncForPC(pc).Name()
	}

	return meta
}
