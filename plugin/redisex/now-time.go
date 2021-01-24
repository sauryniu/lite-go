package redisex

import "github.com/ahl5esoft/lite-go/timeex"

type nowTime struct {
	IRedis
}

func (m nowTime) Unix() int64 {
	t, err := m.Time()
	if err != nil {
		return 0
	}

	return t.Unix()
}

// UnixNano is 当前Unix纳秒级
func (m nowTime) UnixNano() int64 {
	t, err := m.Time()
	if err != nil {
		return 0
	}

	return t.UnixNano()
}

// NewNowTime is redis当前时间
func NewNowTime(redis IRedis) timeex.INowTime {
	return &nowTime{redis}
}
