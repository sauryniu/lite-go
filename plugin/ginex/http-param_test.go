package ginex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testHTTPParamContext struct {
	Ctx      *gin.Context
	Endpoint string
	Name     string
}

func (m testHTTPParamContext) GetGinCtx() *gin.Context {
	return m.Ctx
}

func (m *testHTTPParamContext) SetEndpoint(v string) {
	m.Endpoint = v
}

func (m *testHTTPParamContext) SetAPIName(v string) {
	m.Name = v
}

func Test_NewHTTPParamHandler(t *testing.T) {
	endpoint := "endpoint"
	name := "name"
	resp := httptest.NewRecorder()
	app := gin.New()
	app.POST("/:endpoint/:name", func(ctx *gin.Context) {
		hCtx := &testHTTPParamContext{
			Ctx: ctx,
		}
		err := NewHTTPParamHandler().Handle(hCtx)
		assert.NoError(t, err)
		assert.Equal(t, hCtx.Endpoint, endpoint)
		assert.Equal(t, hCtx.Name, name)
	})
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	app.ServeHTTP(resp, req)
}
