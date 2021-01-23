package timeex

// NowUnixIoCKey is INowUnix依赖注入键
const NowUnixIoCKey = "now-unix"

// INowUnix is 当前unix秒级
type INowUnix interface {
	Unix() int64
}
