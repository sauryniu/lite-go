package cor

// IHandler is 责任链处理器
type IHandler interface {
	Break()
	Handle() error
	IsBreak() bool
	SetDelayNext(NewFunc) IHandler
	SetNext(IHandler) IHandler
}
