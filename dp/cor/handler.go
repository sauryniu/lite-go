package cor

type handler struct {
	HandleFunc func(ctx interface{}) error

	next IHandler
}

func (m handler) Handle(ctx interface{}) (err error) {
	if err = m.HandleFunc(ctx); err != nil || m.next == nil {
		return
	}

	if breakable, ok := ctx.(IBreakable); ok && breakable.IsBreak() {
		return
	}

	return m.next.Handle(ctx)
}

func (m *handler) SetNext(handler IHandler) IHandler {
	m.next = handler
	return handler
}

// New is 创建IHandler
func New(handleFunc func(ctx interface{}) error) IHandler {
	return &handler{
		HandleFunc: handleFunc,
	}
}
