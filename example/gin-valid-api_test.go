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

type ginValidAPIContext struct {
	Resp *httptest.ResponseRecorder
}

type verifyAPI struct {
	*ginex.API

	Name string `binding:"min=1,max=5"`
}

func (m *ginValidAPIContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "verify"
	api.Register(endpoint, name, func() api.IAPI {
		a := &verifyAPI{
			API: &ginex.API{},
		}
		a.API.Derived = a
		return a
	})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader("{}"),
	)
	req.Header.Add("Content-Type", "application/json")
	return m.Resp, req
}

func (m ginValidAPIContext) GetGinMode() string {
	return ""
}

func (m ginValidAPIContext) GetGinPort() int {
	return 0
}

func Test_GinValidAPI(t *testing.T) {
	ctx := new(ginValidAPIContext)
	err := ginex.NewStartupHandler().Handle(ctx)
	assert.NoError(t, err)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "",
		"err":  int(api.VerifyErrorCode),
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
