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

type startupVerifyContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *startupVerifyContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "verify"
	api.Register(endpoint, name, verifyAPI{})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader("{}"),
	)
	req.Header.Add("Content-Type", "application/json")
	return m.Resp, req
}

func (m startupVerifyContext) GetGinMode() string {
	return ""
}

func (m startupVerifyContext) GetGinPort() int {
	return 0
}

func (m startupVerifyContext) HandleGinCtx(ctx *gin.Context) {
	handleGinContet(ctx)
}

type verifyAPI struct {
	Name string `binding:"min=1,max=5"`
}

func (m verifyAPI) Call() (interface{}, error) {
	return nil, nil
}

func (m verifyAPI) GetScope() int {
	return 0
}

func Test_Verify(t *testing.T) {
	ctx := new(startupVerifyContext)
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
		"err":  int(api.VerifyErrorCode),
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
