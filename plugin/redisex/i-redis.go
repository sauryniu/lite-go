package redisex

import "time"

// IoCKey is 依赖注入键
const IoCKey = "redis"

// IRedis is redis接口
type IRedis interface {
	Close() error
	Del(...string) (int, error)
	Exists(string) (bool, error)
	Eval(string, []string, ...interface{}) (interface{}, error)
	Get(string) (string, error)
	Publish(string, interface{}) (int, error)
	Set(string, string, ...interface{}) (bool, error)
	Subscribe(channels []string, handleAction func(sub interface{}))
	Time() (time.Time, error)
}
