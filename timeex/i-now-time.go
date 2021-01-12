package timeex

// NowTimeIoCKey is 依赖注入键
const NowTimeIoCKey = "now-time"

// INowTime is 当前时间接口
type INowTime interface {
	NanoUnix() int64
	Unix() int64
}
