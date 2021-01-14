package ioos

import (
	"io/ioutil"
	"os"

	"github.com/ahl5esoft/lite-go/ioex"
	"github.com/ahl5esoft/lite-go/ioex/iopath"

	underscore "github.com/ahl5esoft/golang-underscore"
)

type directory struct {
	ioex.INode
}

func (m directory) Create() error {
	if m.IsExist() {
		return nil
	}

	dirPath := m.GetPath()
	return os.Mkdir(dirPath, os.ModePerm)
}

func (m directory) FindDirectories() []ioex.IDirectory {
	children := make([]ioex.IDirectory, 0)
	return m.findNodes(children, func(r os.FileInfo, filePath string) interface{} {
		if r.IsDir() {
			children = append(
				children,
				NewDirectory(filePath),
			)
		}
		return children
	}).([]ioex.IDirectory)
}

func (m directory) FindFiles() []ioex.IFile {
	children := make([]ioex.IFile, 0)
	return m.findNodes(children, func(r os.FileInfo, filePath string) interface{} {
		if !r.IsDir() {
			children = append(
				children,
				NewFile(filePath),
			)
		}
		return children
	}).([]ioex.IFile)
}

func (m directory) findNodes(defaultValue interface{}, handleNodeFunc func(r os.FileInfo, nodePath string) interface{}) interface{} {
	dirPath := m.GetPath()
	nodes, err := ioutil.ReadDir(dirPath)
	if err != nil || len(nodes) == 0 {
		return defaultValue
	}

	var res interface{}
	underscore.Chain(nodes).Each(func(r os.FileInfo, _ int) {
		nodePath := iopath.Join(
			dirPath,
			r.Name(),
		)
		res = handleNodeFunc(r, nodePath)
	})
	return res
}

// NewDirectory is 创建ioex.IDirectory
func NewDirectory(pathArgs ...string) ioex.IDirectory {
	return &directory{
		INode: newNode(pathArgs...),
	}
}
