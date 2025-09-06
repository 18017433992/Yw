package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"example.com/blog/response"
	"github.com/gin-gonic/gin"
)

// ErrHandler 全局异常处理中间件
func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录panic信息和堆栈跟踪
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]
				log.Printf("Panic recovered: %v\n%s", err, stack)

				// 判断错误类型并返回相应的响应
				switch e := err.(type) {
				case string:
					response.InternalServerError(c, fmt.Sprintf("服务器内部错误: %s", e))
				case error:
					response.InternalServerError(c, fmt.Sprintf("服务器内部错误: %s", e.Error()))
				default:
					response.InternalServerError(c, "服务器内部错误")
				}

				// 终止请求处理
				c.Abort()
				return
			}

			// 检查响应状态码，如果是错误状态码且没有响应体，则返回统一错误格式
			if c.Writer.Status() >= http.StatusBadRequest && c.Writer.Size() <= 0 {
				switch c.Writer.Status() {
				case http.StatusNotFound:
					response.Error(c, http.StatusNotFound, "请求的资源不存在")
				case http.StatusMethodNotAllowed:
					response.Error(c, http.StatusMethodNotAllowed, "请求方法不被允许")
				case http.StatusInternalServerError:
					response.InternalServerError(c, "服务器内部错误")
				default:
					response.Error(c, c.Writer.Status(), "请求处理失败")
				}
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
