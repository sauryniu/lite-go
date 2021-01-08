package api

// IAPI is API接口
type IAPI interface {
	Auth() bool
	Call() interface{}
	Valid(ctx interface{}) bool
}
