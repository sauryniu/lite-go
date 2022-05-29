package contract

// api会话
type IApiSession interface {
	// 初始化会话
	InitSession(req interface{}) error
}
