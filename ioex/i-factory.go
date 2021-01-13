package ioex

// IFactory is 工厂接口
type IFactory interface {
	BuildDirectory(pathArgs ...string) IDirectory
	BuildFile(pathArgs ...string) IFile
}
