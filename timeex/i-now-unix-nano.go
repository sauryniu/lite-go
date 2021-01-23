package timeex

// NowUnixNanoIoCKey is INowUnixNano依赖注入键
const NowUnixNanoIoCKey = "now-unix-nano"

// INowUnixNano is 当前unix纳秒级
type INowUnixNano interface {
	UnixNano() int64
}
