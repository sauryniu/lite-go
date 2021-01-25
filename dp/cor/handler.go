package cor

import "github.com/ahl5esoft/lite-go/dp/ioc"

type handler struct {
	isBreak bool
	next    IHandler
}

func (m *handler) Break() {
	m.isBreak = true
}

func (m handler) Handle() (err error) {
	if m.next == nil || m.isBreak {
		return
	}

	return m.next.Handle()
}

func (m *handler) SetNext(handler IHandler) IHandler {
	m.next = handler
	ioc.Inject(handler)
	return handler
}

// New is 创建IHandler基类
func New() IHandler {
	return new(handler)
}
