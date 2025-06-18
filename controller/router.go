package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller 接口定义所有 controller 必须实现的方法
type Controller interface {
	RegisterRouter(r gin.IRouter)
}

// RegisterRouters 注册所有路由
func (c *Controllers) RegisterRouters(r gin.IRouter) {
	r.Use(CORSMiddleware())
	api := r.Group("/api")
	c.Product.RegisterRouter(api)
	// ... 注册其他 controller 的路由
}

// Controllers 存储所有的 controller
type Controllers struct {
	Product *Product

	// ... 添加其他 controller
}

func NewGinServer() *gin.Engine {
	return gin.Default()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置 CORS 头
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		// 对于 OPTIONS 请求，直接返回 204
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
