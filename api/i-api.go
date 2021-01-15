package api

// IAPI is api接口
type IAPI interface {
	Call() (interface{}, error)
}
