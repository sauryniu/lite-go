package ginapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type defaultAPI struct{}

func (m defaultAPI) Call() (interface{}, error) {
	return "ok", nil
}

func Test_Default(t *testing.T) {
	endpoint := "endpoint"
	name := "default"
	api.Register(endpoint, name, defaultAPI{})

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	resp := httptest.NewRecorder()
	ginex.Run(
		gin.ReleaseMode,
		ginex.NewPostRunOption(),
		ginex.NewServeHTTPRunOption(req, resp),
	)

	assert.JSONEq(
		t,
		resp.Body.String(),
		`{"data":"ok","err":0}`,
	)
}
