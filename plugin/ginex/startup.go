package ginex

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/gin-gonic/gin"
)

// StartupContext is 启动上下文
type StartupContext struct {
	GinPort int
	GinMode string
}

// NewStartup is 启动处理器
func NewStartup() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(StartupContext); ok {
			if sCtx.GinMode == "" {
				sCtx.GinMode = gin.DebugMode
			}
			gin.SetMode(sCtx.GinMode)

			addr := fmt.Sprintf(":%d", sCtx.GinPort)
			fmt.Println(addr)
			newEngine().Run(addr)
		}
		return nil
	})
}
