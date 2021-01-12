package ginex

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testHTTPVerifyAPI struct {
	Name string `binding:"min=1,max=5"`
}

func (m testHTTPVerifyAPI) Call() (interface{}, error) {
	return nil, nil
}

func (m testHTTPVerifyAPI) GetScope() int {
	return 0
}

type testHTTPVerifyContext struct {
	Ctx *gin.Context

	apiInstance api.IAPI
}

func (m *testHTTPVerifyContext) GetAPI() api.IAPI {
	if m.apiInstance == nil {
		m.apiInstance = new(testHTTPVerifyAPI)
	}

	return m.apiInstance
}

func (m testHTTPVerifyContext) GetGinCtx() *gin.Context {
	return m.Ctx
}

func Test_NewHTTPVerifyHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	app := gin.New()
	app.POST("/", func(ctx *gin.Context) {
		hCtx := &testHTTPVerifyContext{
			Ctx: ctx,
		}
		err := NewHTTPVerifyHandler().Handle(hCtx)
		assert.NoError(t, err)
		assert.Equal(
			t,
			hCtx.GetAPI().(*testHTTPVerifyAPI).Name,
			"hello",
		)
	})
	req, _ := http.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(`{"name":"hello"}`),
	)
	req.Header.Add("Content-Type", "application/json")
	app.ServeHTTP(resp, req)
}
