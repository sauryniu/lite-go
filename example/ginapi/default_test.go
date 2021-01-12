package ginapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/plugin/db/mongodb"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/ahl5esoft/lite-go/plugin/redisex/goredis"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type defaultAPI struct{}

func (m defaultAPI) Call() (interface{}, error) {
	return "ok", nil
}

func (m defaultAPI) GetScope() int {
	return 0
}

type startupContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *startupContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "default"
	api.Register(endpoint, name, defaultAPI{})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m startupContext) GetGinMode() string {
	return ""
}

func (m startupContext) GetGinPort() int {
	return 0
}

func (m startupContext) HandleGinCtx(ctx *gin.Context) {
	handleGinContet(ctx)
}

func Test_Default(t *testing.T) {
	ctx := new(startupContext)
	handler := mongodb.NewStartupHandler()
	handler.SetNext(
		goredis.NewStartupHandler(),
	).SetNext(
		ginex.NewStartupHandler(),
	)
	err := handler.Handle(ctx)
	assert.NoError(t, err)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "ok",
		"err":  0,
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
