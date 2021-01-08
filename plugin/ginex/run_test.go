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

type testAuthAPI struct {
}

func (m testAuthAPI) Auth() bool {
	return false
}

func (m testAuthAPI) Call() interface{} {
	return nil
}

func (m testAuthAPI) Valid(ctx interface{}) bool {
	return false
}

func Test_newEngine_认证(t *testing.T) {
	endpoint := "endpoint"
	name := "auth"
	api.Register(endpoint, name, testAuthAPI{})

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

func (m testValidAPI) Call() interface{} {
	return nil
}

func (m testValidAPI) Valid(ctx interface{}) bool {
	return false
}

func Test_newEngine_验证(t *testing.T) {
	endpoint := "endpoint"
	name := "verify"
	api.Register(endpoint, name, testValidAPI{})

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

type testAPIErrorAPI struct {
}

func (m testAPIErrorAPI) Auth() bool {
	return true
}

func (m testAPIErrorAPI) Call() interface{} {
	api.Throw(11, "msg")
	return nil
}

func (m testAPIErrorAPI) Valid(ctx interface{}) bool {
	return true
}

func Test_newEngine_apierror(t *testing.T) {
	endpoint := "endpoint"
	name := "apierror"
	api.Register(endpoint, name, testAPIErrorAPI{})

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

type testPanicAPI struct {
}

func (m testPanicAPI) Auth() bool {
	return true
}

func (m testPanicAPI) Call() interface{} {
	panic("p")
}

func (m testPanicAPI) Valid(ctx interface{}) bool {
	return true
}

func Test_newEngine_panicAPI(t *testing.T) {
	endpoint := "endpoint"
	name := "panic"
	api.Register(endpoint, name, testPanicAPI{})

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
	publicAPI

	Name string `binding:"min=1,max=10"`
}

func (m testOKAPI) Call() interface{} {
	return m.Name
}

func (m *testOKAPI) Valid(ctx interface{}) bool {
	return m.publicAPI.ValidDerive(ctx, m)
}

func Test_newEngine_ok(t *testing.T) {
	endpoint := "endpoint"
	name := "ok"
	api.Register(
		endpoint,
		name,
		new(testOKAPI),
	)

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
