package redisex

import "github.com/ahl5esoft/lite-go/timeex"

type redisNowTime struct {
	Redis IRedis
}

func (m redisNowTime) NanoUnix() int64 {
	t, err := m.Redis.Time()
	if err != nil {
		return 0
	}

	return t.UnixNano()
}

func (m redisNowTime) Unix() int64 {
	t, err := m.Redis.Time()
	if err != nil {
		return 0
	}

	return t.Unix()
}

// NewNowTime is timeex.INowTime实例
func NewNowTime(redis IRedis) timeex.INowTime {
	return &redisNowTime{
		Redis: redis,
	}
}
