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

// NewTip is 提醒错误
func NewTip(err error) Custom {
	return New(
		TipCode,
		err.Error(),
	)
}

// NewTipf is 提醒错误
func NewTipf(format string, args ...interface{}) Custom {
	return New(TipCode, format, args...)
}

// Throw is 抛出异常
func Throw(code Code, format string, args ...interface{}) {
	panic(
		New(code, format, args...),
	)
}
