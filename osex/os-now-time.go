package osex

import (
	"time"

	"github.com/ahl5esoft/lite-go/timeex"
)

type osNowTime struct{}

func (m osNowTime) Unix() int64 {
	return time.Now().Unix()
}

func (m osNowTime) UnixNano() int64 {
	return time.Now().UnixNano()
}

// NewNowUnix is 系统当前Unix秒级
func NewNowUnix() timeex.INowUnix {
	return new(osNowTime)
}

// NewNowUnixNano is 系统当前Unix纳秒级
func NewNowUnixNano() timeex.INowUnixNano {
	return new(osNowTime)
}
