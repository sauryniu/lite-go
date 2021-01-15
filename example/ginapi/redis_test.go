package ginapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/ahl5esoft/lite-go/plugin/redisex/goredis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

var callTime time.Time

type redisAPI struct {
	Redis redisex.IRedis `inject:"redis"`
}

func (m redisAPI) Call() (interface{}, error) {
	var err error
	callTime, err = m.Redis.Time()
	return callTime, err
}

func Test_Regis(t *testing.T) {
	redisInstance := goredis.New(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ioc.Set(redisex.IoCKey, redisInstance)

	endpoint := "endpoint"
	name := "redis"
	api.Register(endpoint, name, redisAPI{})

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
		Data:  callTime,
		Error: 0,
	})
	assert.JSONEq(
		t,
		res,
		resp.Body.String(),
	)
}
