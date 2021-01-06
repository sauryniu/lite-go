package ioos

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ahl5esoft/lite-go/lib/ioex"
	"github.com/ahl5esoft/lite-go/lib/ioex/iopath"

	"github.com/stretchr/testify/assert"
)

func Test_Build_目录(t *testing.T) {
	_, isDir := Build("a", "b").(ioex.IDirectory)
	assert.True(t, isDir)
}

func Test_Build_文件(t *testing.T) {
	_, isFile := Build("a", "b.txt").(ioex.IFile)
	assert.True(t, isFile)
}

func Test_node_GetName_Dir(t *testing.T) {
	res := newNode("a", "b", "c").GetName()
	assert.Equal(t, res, "c")
}

func Test_node_GetName_File(t *testing.T) {
	res := newNode("a", "b", "c.txt").GetName()
	assert.Equal(t, res, "c.txt")
}

func Test_node_GetParent_Directory(t *testing.T) {
	res := NewDirectory("a", "b").GetParent()
	assert.Equal(
		t,
		res.GetPath(),
		"a",
	)
}

func Test_node_GetParent_File(t *testing.T) {
	res := NewDirectory("a", "b.txt").GetParent()
	assert.Equal(
		t,
		res.GetPath(),
		"a",
	)
}

func Test_node_IsExist_F(t *testing.T) {
	res := newNode("a", "b", "c").IsExist()
	assert.False(t, res)
}

func Test_node_IsExist_T(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	res := newNode(wd, "node.go").IsExist()
	assert.True(t, res)
}

func Test_node_Move_File(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	node := newNode(wd, "move.go")
	err = ioutil.WriteFile(
		node.GetPath(),
		make([]byte, 0),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer node.Remove()

	dstNode := newNode(wd, "move-dst.go")
	defer dstNode.Remove()

	err = node.Move(
		dstNode.GetPath(),
	)
	assert.NoError(t, err)

	assert.False(
		t,
		node.IsExist(),
	)
	assert.True(
		t,
		dstNode.IsExist(),
	)
}

func Test_node_Move_FileIsExist(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	node := newNode(wd, "move.go")

	err = ioutil.WriteFile(
		node.GetPath(),
		make([]byte, 0),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer node.Remove()

	dstNode := newNode(wd, "move-dst.go")

	err = ioutil.WriteFile(
		dstNode.GetPath(),
		make([]byte, 0),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer dstNode.Remove()

	err = node.Move(
		dstNode.GetPath(),
	)
	assert.NoError(t, err)

	assert.False(
		t,
		node.IsExist(),
	)
	assert.True(
		t,
		dstNode.IsExist(),
	)
}

func Test_node_Remove_Directory(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	node := newNode(wd, "remove")
	err = os.Mkdir(
		node.GetPath(),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer os.Remove(
		node.GetPath(),
	)

	err = node.Remove()
	assert.NoError(t, err)

	ok := node.IsExist()
	assert.False(t, ok)
}

func Test_node_Remove_DirectoryHasChildren(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	node := newNode(wd, "remove")
	err = os.Mkdir(
		node.GetPath(),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer os.Remove(
		node.GetPath(),
	)

	childNodePath := iopath.Join(
		node.GetPath(),
		"one",
	)
	err = os.Mkdir(
		childNodePath,
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer os.Remove(childNodePath)

	err = node.Remove()
	assert.NoError(t, err)

	ok := node.IsExist()
	assert.False(t, ok)
}

func Test_node_Remove_File(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	node := newNode(wd, "remove.go")
	err = ioutil.WriteFile(
		node.GetPath(),
		make([]byte, 0),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer os.Remove(
		node.GetPath(),
	)

	err = node.Remove()
	assert.NoError(t, err)

	ok := node.IsExist()
	assert.False(t, ok)
}

func Test_node_Remove_不存在(t *testing.T) {
	err := newNode("a", "b", "c").Remove()
	assert.NoError(t, err)
}
