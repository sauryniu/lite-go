package osex

import (
	"os"
	"path/filepath"

	"github.com/ahl5esoft/lite-go/ioex"
)

type ioNode struct {
	ioPath ioex.IPath
	path   string
}

func (m ioNode) GetName() string {
	return filepath.Base(m.path)
}

func (m ioNode) GetParent() ioex.IDirectory {
	return NewIODirectory(
		m.ioPath,
		m.GetPath(),
		"..",
	)
}

func (m ioNode) GetPath() string {
	return m.path
}

func (m ioNode) IsExist() bool {
	_, err := os.Stat(m.path)
	return err == nil || os.IsExist(err)
}

func (m ioNode) Move(paths ...string) error {
	dstPath := m.ioPath.Join(paths...)
	return os.Rename(
		m.GetPath(),
		dstPath,
	)
}

func (m ioNode) Remove() error {
	if !m.IsExist() {
		return nil
	}

	return os.RemoveAll(
		m.GetPath(),
	)
}

func newIONode(ioPath ioex.IPath, paths ...string) ioex.INode {
	return ioNode{
		ioPath: ioPath,
		path:   ioPath.Join(paths...),
	}
}
