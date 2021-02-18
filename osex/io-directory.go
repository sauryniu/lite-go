package osex

import (
	"io/ioutil"
	"os"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/ioex"
)

type ioDirectory struct {
	ioex.INode

	ioPath ioex.IPath
}

func (m ioDirectory) Create() error {
	if m.IsExist() {
		return nil
	}

	dirPath := m.GetPath()
	return os.Mkdir(dirPath, os.ModePerm)
}

func (m ioDirectory) FindDirectories() []ioex.IDirectory {
	children := make([]ioex.IDirectory, 0)
	return m.findNodes(children, func(r os.FileInfo, filePath string) interface{} {
		if r.IsDir() {
			children = append(
				children,
				NewIODirectory(m.ioPath, filePath),
			)
		}
		return children
	}).([]ioex.IDirectory)
}

func (m ioDirectory) FindFiles() []ioex.IFile {
	children := make([]ioex.IFile, 0)
	return m.findNodes(children, func(r os.FileInfo, filePath string) interface{} {
		if !r.IsDir() {
			children = append(
				children,
				NewIOFile(m.ioPath, filePath),
			)
		}
		return children
	}).([]ioex.IFile)
}

func (m ioDirectory) findNodes(defaultValue interface{}, handleNodeFunc func(r os.FileInfo, nodePath string) interface{}) interface{} {
	dirPath := m.GetPath()
	nodes, err := ioutil.ReadDir(dirPath)
	if err != nil || len(nodes) == 0 {
		return defaultValue
	}

	var res interface{}
	underscore.Chain(nodes).Each(func(r os.FileInfo, _ int) {
		nodePath := m.ioPath.Join(
			dirPath,
			r.Name(),
		)
		res = handleNodeFunc(r, nodePath)
	})
	return res
}

// NewIODirectory is 创建ioex.IDirectory
func NewIODirectory(ioPath ioex.IPath, paths ...string) ioex.IDirectory {
	return &ioDirectory{
		INode:  newIONode(ioPath, paths...),
		ioPath: ioPath,
	}
}
