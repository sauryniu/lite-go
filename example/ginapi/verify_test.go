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
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type verifyAPI struct {
	Name string `binding:"min=1,max=5,required" json:"na"`
}

func (m verifyAPI) Call() (interface{}, error) {
	return m.Name, nil
}

func Test_Verify_OK(t *testing.T) {
	endpoint := "endpoint"
	name := "verify"
	api.Register(endpoint, name, verifyAPI{})

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(`{"na":"go"}`),
	)
	req.Header.Add("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	ginex.Run(
		gin.ReleaseMode,
		ginex.NewPostRunOption(),
		ginex.NewServeHTTPRunOption(req, resp),
	)

	res, _ := jsoniter.MarshalToString(api.Response{
		Data:  "go",
		Error: 0,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}

func Test_Verify_Fail(t *testing.T) {
	endpoint := "endpoint"
	name := "verify"
	api.Register(endpoint, name, verifyAPI{})

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(`{}`),
	)
	req.Header.Add("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	ginex.Run(
		gin.ReleaseMode,
		ginex.NewPostRunOption(),
		ginex.NewServeHTTPRunOption(req, resp),
	)

	res, _ := jsoniter.MarshalToString(api.Response{
		Data:  "",
		Error: api.VerifyErrorCode,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}
