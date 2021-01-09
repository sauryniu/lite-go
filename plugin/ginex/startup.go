package ginex

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/gin-gonic/gin"
)

// IStartupContext is 启动上下文接口
type IStartupContext interface {
	GetGinPort() int
	GetGinMode() string
}

// NewStartupHandler is 启动处理器
func NewStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(IStartupContext); ok {
			mode := sCtx.GetGinMode()
			if mode == "" {
				mode = gin.DebugMode
			}
			gin.SetMode(mode)

			addr := fmt.Sprintf(
				":%d",
				sCtx.GetGinPort(),
			)
			fmt.Println(addr)
			newEngine().Run(addr)
		}
		return nil
	})
}
