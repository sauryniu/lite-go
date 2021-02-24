//go:generate mockgen -destination i-redis_mock.go -package redisex github.com/ahl5esoft/lite-go/plugin/redisex IRedis

package redisex

import (
	"time"

	"github.com/ahl5esoft/lite-go/plugin/pubsub"
)

// IRedis is redis接口
type IRedis interface {
	pubsub.IPublisher
	pubsub.ISubscriber

	Close() error
	Del(...string) (int, error)
	Exists(string) (bool, error)
	Eval(string, []string, ...interface{}) (interface{}, error)
	Get(string) (string, error)
	Set(string, string, ...interface{}) (bool, error)
	Time() (time.Time, error)
	TTL(key string) (time.Duration, error)
}
