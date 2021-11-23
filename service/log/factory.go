package log

import (
	"github.com/ahl5esoft/lite-go/contract"
)

type BuildFunc func() contract.ILog

type factory BuildFunc

func (m factory) Build() contract.ILog {
	return m()
}

func NewFactory(build BuildFunc) contract.ILogFactory {
	return factory(build)
}
