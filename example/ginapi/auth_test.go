package ginapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/dp/cor"
	"github.com/ahl5esoft/lite-go/plugin/db/mongodb"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/ahl5esoft/lite-go/plugin/redisex/goredis"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type authStartupContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *authStartupContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "auth"
	api.Register(endpoint, name, defaultAPI{})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m authStartupContext) GetGinMode() string {
	return ""
}

func (m authStartupContext) GetGinPort() int {
	return 0
}

func (m authStartupContext) HandleGinCtx(ctx *gin.Context) {
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
		cor.New(func(_ interface{}) error {
			return api.NewError(api.AuthErrorCode, "")
		}),
	).SetNext(
		ginex.NewHTTPVerifyHandler(),
	)
	err = handler.Handle(hCtx)
}

func Test_Auth(t *testing.T) {
	ctx := new(authStartupContext)
	handler := mongodb.NewStartupHandler()
	handler.SetNext(
		goredis.NewStartupHandler(),
	).SetNext(
		ginex.NewStartupHandler(),
	)
	err := handler.Handle(ctx)
	assert.NoError(t, err)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "",
		"err":  int(api.AuthErrorCode),
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
