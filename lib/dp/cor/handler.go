package cor

type handler struct {
	HandleFunc func(ctx interface{}) error

	next IHandler
}

func (m handler) Handle(ctx interface{}) error {
	if err := m.HandleFunc(ctx); err != nil {
		return err
	}

	if m.next != nil {
		return m.next.Handle(ctx)
	}

	return nil
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
