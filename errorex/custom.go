package errorex

import "fmt"

// Custom is 自定义
type Custom struct {
	error

	Code Code
}

// New is 实例化错误结构
func New(code Code, format string, args ...interface{}) Custom {
	return Custom{
		fmt.Errorf(format, args...),
		code,
	}
}

// Throw is 抛出异常
func Throw(code Code, format string, args ...interface{}) {
	panic(
		New(code, format, args...),
	)
}
