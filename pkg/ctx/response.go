package ctx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

type Response struct {
	Code    int         `json:"code"`           // 状态码
	Message string      `json:"message"`        // 提示信息
	Data    interface{} `json:"data,omitempty"` // 返回数据
	Meta    any         `json:"metadata,omitempty"`
}

func Success(c *gin.Context, data interface{}) error {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
	return nil
}

func Error(c *gin.Context, code int, message string) error {

	resp := Response{
		Code:    code,
		Message: message,
		Data:    nil,
		Meta:    initMeta(),
	}
	if resp.Data == nil {
		resp.Data = ""
	}

	c.AbortWithStatusJSON(code, resp)
	return nil
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

func HandlerFunc(fn func(ctx *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Meta:    initMeta(),
			})
			return
		}
	}
}
