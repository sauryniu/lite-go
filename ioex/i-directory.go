package ioex

// IDirectory is 目录接口
type IDirectory interface {
	INode

	Create() error
	FindDirectories() []IDirectory
	FindFiles() []IFile
}
