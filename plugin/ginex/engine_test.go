package ginex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func Test_newEngine_认证(t *testing.T) {
	endpoint := "endpoint"
	name := "auth"
	api.Register(endpoint, name, func() api.IAPI {
		return &API{
			AuthFunc: func() bool {
				return false
			},
		}
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	newEngine().ServeHTTP(resp, req)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "",
		"err":  api.AuthErrorCode,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}

type testValidAPI struct {
}

func (m testValidAPI) Auth() bool {
	return true
}

func (m testValidAPI) Call() (interface{}, error) {
	return nil, nil
}

func (m testValidAPI) SetRequest(_ interface{}) {

}

func (m testValidAPI) Valid() bool {
	return false
}

func Test_newEngine_验证(t *testing.T) {
	endpoint := "endpoint"
	name := "verify"
	api.Register(endpoint, name, func() api.IAPI {
		return &testValidAPI{}
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	newEngine().ServeHTTP(resp, req)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "",
		"err":  api.VerifyErrorCode,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}

func Test_newEngine_InvalidAPI(t *testing.T) {
	endpoint := "endpoint"
	name := "invalid"

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	newEngine().ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "",
		"err":  api.APIErrorCode,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}

func Test_newEngine_apierror(t *testing.T) {
	endpoint := "endpoint"
	name := "apierror"
	api.Register(endpoint, name, func() api.IAPI {
		return &API{
			CallFunc: func() (interface{}, error) {
				api.Throw(11, "msg")
				return nil, nil
			},
		}
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	newEngine().ServeHTTP(resp, req)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "msg",
		"err":  11,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}

func Test_newEngine_panicAPI(t *testing.T) {
	endpoint := "endpoint"
	name := "panic"
	api.Register(endpoint, name, func() api.IAPI {
		return &API{
			CallFunc: func() (interface{}, error) {
				panic("p")
			},
		}
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	newEngine().ServeHTTP(resp, req)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "",
		"err":  599,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}

type testOKAPI struct {
	*API

	Name string `binding:"min=1,max=10"`
}

func Test_newEngine_ok(t *testing.T) {
	endpoint := "endpoint"
	name := "ok"
	api.Register(endpoint, name, func() api.IAPI {
		var a *testOKAPI
		a = &testOKAPI{
			API: &API{
				CallFunc: func() (interface{}, error) {
					return "hello", nil
				},
			},
		}
		return a
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(`{"name":"hello"}`),
	)
	req.Header.Add("Content-Type", "application/json")
	newEngine().ServeHTTP(resp, req)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": "hello",
		"err":  0,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}
