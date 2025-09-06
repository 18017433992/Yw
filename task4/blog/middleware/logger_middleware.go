package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LogRequestMiddleware 是一个日志中间件，用于记录请求信息
func LogRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 客户端IP
		clientIP := c.ClientIP()

		// 请求方式
		method := c.Request.Method

		// 请求的路由
		path := c.Request.URL.Path

		// 状态码
		status := c.Writer.Status()

		// 记录日志信息
		fmt.Printf("[%s] [%s] [%s] [%s] [%d] [%v]\n", end.Format("2006-01-02 15:04:05"), clientIP, method, path, status, latency)
	}
}
