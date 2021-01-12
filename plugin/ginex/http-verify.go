package ginex

import (
	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/gin-gonic/gin"
)

// IHTTPVerifyContext is http验证上下文
type IHTTPVerifyContext interface {
	GetAPI() api.IAPI
	GetGinCtx() *gin.Context
}

// NewHTTPVerifyHandler is http验证处理
func NewHTTPVerifyHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if rCtx, ok := ctx.(IHTTPVerifyContext); ok {
			apiInstance := rCtx.GetAPI()
			if apiInstance != nil {
				if err := rCtx.GetGinCtx().ShouldBind(apiInstance); err != nil {
					return api.NewError(api.VerifyErrorCode, "")
				}

				return nil
			}
		}
		return api.NewError(api.APIErrorCode, "")
	})
}
