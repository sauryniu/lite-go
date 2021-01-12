package ginapi

import (
	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/gin-gonic/gin"
)

type httpContext struct {
	Ctx *gin.Context
	Err error

	api      api.IAPI
	endpoint string
	apiName  string
}

func (m httpContext) GetAPI() api.IAPI {
	return m.api
}

func (m httpContext) GetAPIName() string {
	return m.apiName
}

func (m httpContext) GetEndpoint() string {
	return m.endpoint
}

func (m httpContext) GetGinCtx() *gin.Context {
	return m.Ctx
}

func (m *httpContext) SetAPI(v api.IAPI) {
	m.api = v
}

func (m *httpContext) SetAPIName(v string) {
	m.apiName = v
}

func (m *httpContext) SetEndpoint(v string) {
	m.endpoint = v
}

func handleGinContet(ctx *gin.Context) {
	var err error
	hCtx := &httpContext{
		Ctx: ctx,
	}
	defer func() {
		ginex.JSON(
			err,
			hCtx.GetAPI(),
			ctx,
		)
	}()

	handler := ginex.NewHTTPParamHandler()
	handler.SetNext(
		api.NewCreateHandler(),
	).SetNext(
		ginex.NewHTTPVerifyHandler(),
	)
	err = handler.Handle(hCtx)
}
