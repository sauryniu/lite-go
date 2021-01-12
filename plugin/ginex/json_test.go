package ginex

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func Test_JSON(t *testing.T) {
	resp := httptest.NewRecorder()
	app := gin.New()
	app.POST("/", func(ctx *gin.Context) {
		JSON(nil, nil, ctx)

		res, _ := jsoniter.MarshalToString(jsonResponse{
			Data:  "",
			Error: api.APIErrorCode,
		})
		assert.Equal(
			t,
			resp.Body.String(),
			res,
		)
	})
	req, _ := http.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(""),
	)
	app.ServeHTTP(resp, req)
}
