package ginapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/plugin/db/mongodb"
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

func (m redisAPI) GetScope() int {
	return 0
}

type startupRedisContext struct {
	Resp *httptest.ResponseRecorder
}

func (m *startupRedisContext) GetGinMock() (http.ResponseWriter, *http.Request) {
	endpoint := "endpoint"
	name := "redis"
	api.Register(endpoint, name, redisAPI{})

	m.Resp = httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/%s/%s", endpoint, name),
		strings.NewReader(""),
	)
	return m.Resp, req
}

func (m startupRedisContext) GetGinMode() string {
	return ""
}

func (m startupRedisContext) GetGinPort() int {
	return 0
}

func (m startupRedisContext) GetRedisOption() *redis.Options {
	return &redis.Options{
		Addr: "127.0.0.1:6379",
	}
}

func (m startupRedisContext) GetRedisLockOption() *redis.Options {
	return nil
}

func (m startupRedisContext) HandleGinCtx(ctx *gin.Context) {
	handleGinContet(ctx)
}

func Test_Regis(t *testing.T) {
	ctx := new(startupRedisContext)
	handler := mongodb.NewStartupHandler()
	handler.SetNext(
		goredis.NewStartupHandler(),
	).SetNext(
		ginex.NewStartupHandler(),
	)
	err := handler.Handle(ctx)
	assert.NoError(t, err)

	res, _ := jsoniter.MarshalToString(map[string]interface{}{
		"data": callTime,
		"err":  0,
	})
	assert.JSONEq(
		t,
		res,
		ctx.Resp.Body.String(),
	)
}
