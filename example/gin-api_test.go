package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type ginAPIContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *ginAPIContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "default"
	api.Register(endpoint, name, func() api.IAPI {
		return &ginex.API{
			CallFunc: func() (interface{}, error) {
				return "ok", nil
			},
		}
	})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m ginAPIContext) GetGinMode() string {
	return ""
}

func (m ginAPIContext) GetGinPort() int {
	return 0
}

func Test_GinAPI(t *testing.T) {
	ctx := new(ginAPIContext)
	err := ginex.NewStartupHandler().Handle(ctx)
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
