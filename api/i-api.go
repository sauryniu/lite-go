package api

// IAPI is api接口
type IAPI interface {
	Call() (interface{}, error)
}

// IAPIScope is api作用域接口
type IAPIScope interface {
	GetScope() int
}
