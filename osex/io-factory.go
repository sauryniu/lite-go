package osex

import "github.com/ahl5esoft/lite-go/ioex"

type ioFactory struct {
	ioPath ioex.IPath
}

func (m ioFactory) BuildDirectory(paths ...string) ioex.IDirectory {
	return NewIODirectory(m.ioPath, paths...)
}

func (m ioFactory) BuildFile(paths ...string) ioex.IFile {
	return NewIOFile(m.ioPath, paths...)
}

// NewIOFactory is 创建ioex.IFactory实例
func NewIOFactory(ioPath ioex.IPath) ioex.IFactory {
	return &ioFactory{
		ioPath: ioPath,
	}
}
