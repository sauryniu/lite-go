package osex

import (
	"time"
)

// NowTime is 当前时间
type NowTime struct{}

// Unix is Unix秒级
func (m NowTime) Unix() int64 {
	return time.Now().Unix()
}

// UnixNano is Unix纳秒级
func (m NowTime) UnixNano() int64 {
	return time.Now().UnixNano()
}
