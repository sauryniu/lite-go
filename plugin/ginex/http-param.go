package ginex

import (
	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/gin-gonic/gin"
)

type httpParam struct {
	Endpoint string `uri:"endpoint" binding:"required"`
	Name     string `uri:"name" binding:"required"`
}

// IHTTPParamContext is http参数上下文
type IHTTPParamContext interface {
	GetGinCtx() *gin.Context
	SetAPIName(string)
	SetEndpoint(string)
}

// NewHTTPParamHandler is http参数处理
func NewHTTPParamHandler() cor.IHandler {
	return cor.New(func(ctx interface{}) error {
		if rCtx, ok := ctx.(IHTTPParamContext); ok {
			var param httpParam
			if err := rCtx.GetGinCtx().ShouldBindUri(&param); err != nil {
				return api.NewError(api.VerifyErrorCode, "")
			}

			rCtx.SetEndpoint(param.Endpoint)
			rCtx.SetAPIName(param.Name)
		}

		return nil
	})
}
