package redisex

import "time"

// IoCKey is 依赖注入键
const IoCKey = "redis"

// IRedis is redis接口
type IRedis interface {
	Del(...string) (int, error)
	Exists(string) (bool, error)
	Eval(string, []string, ...interface{}) (interface{}, error)
	Get(string) (string, error)
	Set(string, string, ...interface{}) (bool, error)
	Time() (time.Time, error)
}
