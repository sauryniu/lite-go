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

type ginPanicAPIContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *ginPanicAPIContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "panic"
	api.Register(endpoint, name, func() api.IAPI {
		return &ginex.API{
			CallFunc: func() (interface{}, error) {
				panic("p")
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

func (m ginPanicAPIContext) GetGinMode() string {
	return ""
}

func (m ginPanicAPIContext) GetGinPort() int {
	return 0
}

func Test_GinPanicAPI(t *testing.T) {
	ctx := new(ginPanicAPIContext)
	err := ginex.NewStartupHandler().Handle(ctx)
	assert.NoError(t, err)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "",
		"err":  int(api.PanicErrorCode),
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
