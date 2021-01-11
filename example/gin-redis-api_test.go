package example

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/plugin/ginex"
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/ahl5esoft/lite-go/plugin/redisex/goredis"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type ginRedisAPIContext struct {
	Resp *httptest.ResponseRecorder
	Time time.Time
}

func (m *ginRedisAPIContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "redis"
	api.Register(endpoint, name, func() api.IAPI {
		var a *redisAPI
		a = &redisAPI{
			API: &ginex.API{
				CallFunc: func() (interface{}, error) {
					var err error
					m.Time, err = a.Redis.Time()
					return m.Time, err
				},
			},
		}
		return a
	})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m ginRedisAPIContext) GetGinMode() string {
	return ""
}

func (m ginRedisAPIContext) GetGinPort() int {
	return 0
}

func (m ginRedisAPIContext) GetRedisOption() *redis.Options {
	return &redis.Options{
		Addr: "127.0.0.1:6379",
	}
}

func (m ginRedisAPIContext) GetRedisLockOption() *redis.Options {
	return nil
}

type redisAPI struct {
	*ginex.API

	Redis redisex.IRedis `inject:"redis"`
}

func Test_GinRegisAPI(t *testing.T) {
	handler := goredis.NewStartupHandler()
	handler.SetNext(
		ginex.NewStartupHandler(),
	)
	ctx := new(ginRedisAPIContext)
	err := handler.Handle(ctx)
	assert.NoError(t, err)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": ctx.Time,
		"err":  0,
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
