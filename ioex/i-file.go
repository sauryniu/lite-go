package ioex

import "os"

// IFile is 文件接口
type IFile interface {
	INode

	GetExt() string
	GetFile() (*os.File, error)
	Read(data interface{}) error
	ReadJSON(data interface{}) error
	ReadYaml(data interface{}) error
	Write(data interface{}) error
}
