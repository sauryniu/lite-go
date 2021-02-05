package cor

import (
	"github.com/ahl5esoft/lite-go/dp/ioc"
)

// NewFunc is 创建IHandler实例
type NewFunc func() IHandler

type handler struct {
	isBreak bool
	nexts   []interface{}
}

func (m *handler) Break() {
	m.isBreak = true
}

func (m handler) Handle() (err error) {
	if m.isBreak {
		return
	}

	for _, r := range m.nexts {
		var h IHandler
		if newFunc, ok := r.(NewFunc); ok {
			h = newFunc()
		} else {
			h = r.(IHandler)
		}
		ioc.Inject(h)
		if err = h.Handle(); err != nil || h.IsBreak() {
			break
		}
	}

	return
}

func (m handler) IsBreak() bool {
	return m.isBreak
}

func (m *handler) SetNext(handler IHandler) IHandler {
	m.nexts = append(m.nexts, handler)
	return m
}

func (m *handler) SetDelayNext(newFunc NewFunc) IHandler {
	m.nexts = append(m.nexts, newFunc)
	return m
}

// New is 创建IHandler基类
func New() IHandler {
	return &handler{
		nexts: make([]interface{}, 0),
	}
}
