package ginex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type defaultAPI struct{}

func (m defaultAPI) Call() (interface{}, error) {
	return "ok", nil
}

func Test_NewPostRunOption(t *testing.T) {
	endpoint := "endpoint"
	name := "default"
	api.Register(endpoint, name, defaultAPI{})

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	resp := httptest.NewRecorder()
	Run(
		gin.ReleaseMode,
		NewPostRunOption(),
		func(app *gin.Engine) {
			app.ServeHTTP(resp, req)
		},
	)
	assert.Equal(
		t,
		resp.Body.String(),
		`{"data":"ok","err":0}`,
	)
}
