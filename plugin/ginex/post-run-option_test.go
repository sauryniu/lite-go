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

func Test_NewPostRunOption_API(t *testing.T) {
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

type jsonBodyAPI struct {
	Name string `validate:"max=6,min=1"`
}

func (m jsonBodyAPI) Call() (interface{}, error) {
	return m.Name, nil
}

func Test_NewPostRunOption_JSONBody(t *testing.T) {
	endpoint := "endpoint"
	name := "json"
	api.Register(endpoint, name, jsonBodyAPI{})

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(`{"name":""}`),
	)
	req.Header.Add("Content-Type", "application/json")
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
		`{"data":"","err":503}`,
	)
}

type jsonBodyWithoutVerifyAPI struct {
	Name string
}

func (m jsonBodyWithoutVerifyAPI) Call() (interface{}, error) {
	return m.Name, nil
}

func Test_NewPostRunOption_jsonBodyWithoutVerify(t *testing.T) {
	endpoint := "endpoint"
	name := "jsonWithoutVerify"
	api.Register(endpoint, name, jsonBodyWithoutVerifyAPI{})

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(`{"name":"hello"}`),
	)
	req.Header.Add("Content-Type", "application/json")
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
		`{"data":"hello","err":0}`,
	)
}
