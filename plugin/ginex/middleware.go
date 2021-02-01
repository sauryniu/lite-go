package ginex

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// NewCORSMiddleware is 跨域中间件
func NewCORSMiddleware(origin string, methods ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if origin == "" {
			origin = "*"
		}

		ctx.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
		ctx.Header(
			"Access-Control-Allow-Methods",
			strings.Join(methods, ","),
		)
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		ctx.Next()
	}
}
