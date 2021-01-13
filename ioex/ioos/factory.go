package ioos

import "github.com/ahl5esoft/lite-go/ioex"

type factory struct{}

func (m factory) BuildDirectory(pathArgs ...string) ioex.IDirectory {
	return NewDirectory(pathArgs...)
}

func (m factory) BuildFile(pathArgs ...string) ioex.IFile {
	return NewFile(pathArgs...)
}

// NewFactory is 系统工厂实例
func NewFactory() ioex.IFactory {
	return new(factory)
}
