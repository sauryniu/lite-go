package ioex

// INode is IO节点接口
type INode interface {
	GetName() string
	GetParent() IDirectory
	GetPath() string
	IsExist() bool
	Move(pathArgs ...string) error
	Remove() error
}
