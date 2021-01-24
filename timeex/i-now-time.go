package timeex

// INowTime is 当前时间接口
type INowTime interface {
	Unix() int64
	UnixNano() int64
}
