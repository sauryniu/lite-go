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

func Test_NewStartupHandler_NilCtx(t *testing.T) {
	err := NewStartupHandler().Handle(nil)
	assert.NoError(t, err)
}

type mockStartupContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *mockStartupContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "name"
	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m mockStartupContext) GetGinMode() string {
	return ""
}

func (m mockStartupContext) GetGinPort() int {
	return 0
}

func (m mockStartupContext) HandleGinCtx(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

func Test_NewStartupHandler_Mock(t *testing.T) {
	ctx := new(mockStartupContext)
	err := NewStartupHandler().Handle(ctx)
	assert.NoError(t, err)
	assert.Equal(
		t,
		ctx.Resp.Body.String(),
		"ok",
	)
}
