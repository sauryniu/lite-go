package ginex

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewServeHTTPRunOption(t *testing.T) {
	req, _ := http.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(""),
	)
	resp := httptest.NewRecorder()
	Run(
		gin.ReleaseMode,
		func(app *gin.Engine) {
			app.POST("/", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "ok")
			})
		},
		NewServeHTTPRunOption(req, resp),
	)
	assert.Equal(
		t,
		resp.Body.String(),
		"ok",
	)
}
