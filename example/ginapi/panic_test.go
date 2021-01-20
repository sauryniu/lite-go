package ginapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type panicAPI struct{}

func (m panicAPI) Call() (interface{}, error) {
	panic("p")
}

func Test_Panic(t *testing.T) {
	endpoint := "endpoint"
	name := "panic"
	api.Register(endpoint, name, panicAPI{})

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

	res, _ := jsoniter.MarshalToString(api.Response{
		Data:  "",
		Error: errorex.PanicCode,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}
