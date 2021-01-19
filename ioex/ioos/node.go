package ioos

import (
	"os"
	"path/filepath"

	"github.com/ahl5esoft/lite-go/ioex"
	"github.com/ahl5esoft/lite-go/ioex/iopath"
)

type node struct {
	path string
}

func (m node) GetName() string {
	return filepath.Base(m.path)
}

func (m node) GetParent() ioex.IDirectory {
	return Build(
		m.GetPath(),
		"..",
	).(ioex.IDirectory)
}

func (m node) GetPath() string {
	return m.path
}

func (m node) IsExist() bool {
	_, err := os.Stat(m.path)
	return err == nil || os.IsExist(err)
}

func (m node) Move(dstPath string) error {
	return os.Rename(
		m.GetPath(),
		dstPath,
	)
}

func (m node) Remove() error {
	if !m.IsExist() {
		return nil
	}

	return os.RemoveAll(
		m.GetPath(),
	)
}

// Build is 创建INode
func Build(pathArgs ...string) ioex.INode {
	path := iopath.Join(pathArgs...)
	ext := filepath.Ext(path)
	if ext == "" {
		return NewDirectory(path)
	}

	return NewFile(path)
}

func newNode(pathArgs ...string) ioex.INode {
	return node{
		path: iopath.Join(pathArgs...),
	}
}
