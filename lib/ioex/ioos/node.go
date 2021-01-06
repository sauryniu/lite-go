package ioos

import (
	"os"
	"path/filepath"

	"github.com/ahl5esoft/lite-go/lib/ioex"
	"github.com/ahl5esoft/lite-go/lib/ioex/iopath"
)

type node struct {
	Path string
}

func (m node) GetName() string {
	return filepath.Base(m.Path)
}

func (m node) GetParent() ioex.IDirectory {
	return Build(
		m.GetPath(),
		"..",
	).(ioex.IDirectory)
}

func (m node) GetPath() string {
	return m.Path
}

func (m node) IsExist() bool {
	_, err := os.Stat(m.Path)
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
		Path: iopath.Join(pathArgs...),
	}
}
