package redisex

// NowTime is 当前时间
type NowTime struct {
	Redis IRedis
}

// Unix is 当前Unix秒级
func (m NowTime) Unix() int64 {
	t, err := m.Redis.Time()
	if err != nil {
		return 0
	}

	return t.Unix()
}

// UnixNano is 当前Unix纳秒级
func (m NowTime) UnixNano() int64 {
	t, err := m.Redis.Time()
	if err != nil {
		return 0
	}

	return t.UnixNano()
}
