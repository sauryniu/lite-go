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

type ginAuthAPIContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *ginAuthAPIContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "auth"
	api.Register(endpoint, name, func() api.IAPI {
		return &ginex.API{
			AuthFunc: func() bool {
				return false
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

func (m ginAuthAPIContext) GetGinMode() string {
	return ""
}

func (m ginAuthAPIContext) GetGinPort() int {
	return 0
}

func Test_GinAuthAPI(t *testing.T) {
	ctx := new(ginAuthAPIContext)
	err := ginex.NewStartupHandler().Handle(ctx)
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
