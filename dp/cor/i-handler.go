package cor

// IHandler is 责任链处理器
type IHandler interface {
	Handle(ctx interface{}) error
	SetNext(handler IHandler) IHandler
}
