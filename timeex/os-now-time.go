package timeex

import "time"

type osNowTime struct{}

func (m osNowTime) NanoUnix() int64 {
	return time.Now().UnixNano()
}

func (m osNowTime) Unix() int64 {
	return time.Now().Unix()
}

// NewOSNowTime is INowTime实例
func NewOSNowTime() INowTime {
	return new(osNowTime)
}
