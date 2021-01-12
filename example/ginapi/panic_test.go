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

type panicAPI struct{}

func (m panicAPI) Call() (interface{}, error) {
	panic("p")
}

func (m panicAPI) GetScope() int {
	return 0
}

type startupPanicContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *startupPanicContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "panic"
	api.Register(endpoint, name, panicAPI{})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m startupPanicContext) GetGinMode() string {
	return ""
}

func (m startupPanicContext) GetGinPort() int {
	return 0
}

func (m startupPanicContext) HandleGinCtx(ctx *gin.Context) {
	handleGinContet(ctx)
}

func Test_Panic(t *testing.T) {
	ctx := new(startupPanicContext)
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
		"err":  int(api.PanicErrorCode),
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
