package api

import "fmt"

// CustomError is 自定义错误
type CustomError struct {
	error

	Code ErrorCode
}

// NewError is 实例化错误结构
func NewError(errorCode ErrorCode, format string, args ...interface{}) CustomError {
	return CustomError{
		fmt.Errorf(format, args...),
		errorCode,
	}
}

// Throw is 抛出异常
func Throw(errorCode ErrorCode, format string, args ...interface{}) {
	panic(
		NewError(errorCode, format, args...),
	)
}
