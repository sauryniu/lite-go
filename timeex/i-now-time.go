package timeex

// NowTimeIoCKey is INowTime依赖注入键
const NowTimeIoCKey = "now-unix"

// INowTime is 当前时间接口
type INowTime interface {
	Unix() int64
	UnixNano() int64
}
