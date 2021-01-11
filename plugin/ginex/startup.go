package ginex

import (
	"fmt"
	"net/http"

	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/gin-gonic/gin"
)

// IStartupContext is 启动上下文接口
type IStartupContext interface {
	GetGinMock() (http.ResponseWriter, *http.Request)
	GetGinMode() string
	GetGinPort() int
}

// NewStartupHandler is 启动处理器
func NewStartupHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if sCtx, ok := ctx.(IStartupContext); ok {
			engine := newEngine()
			resp, req := sCtx.GetGinMock()
			if resp != nil && req != nil {
				engine.ServeHTTP(resp, req)
			} else {
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
				engine.Run(addr)
			}
		}
		return nil
	})
}
