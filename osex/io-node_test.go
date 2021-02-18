package osex

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ioNode_GetName_Dir(t *testing.T) {
	ioPath := NewIOPath()
	res := newIONode(ioPath, "a", "b", "c").GetName()
	assert.Equal(t, res, "c")
}

func Test_ioNode_GetName_File(t *testing.T) {
	ioPath := NewIOPath()
	res := newIONode(ioPath, "a", "b", "c.txt").GetName()
	assert.Equal(t, res, "c.txt")
}

func Test_ioNode_GetParent_Directory(t *testing.T) {
	ioPath := NewIOPath()
	res := newIONode(ioPath, "a", "b").GetParent()
	assert.Equal(
		t,
		res.GetPath(),
		"a",
	)
}

func Test_ioNode_GetParent_File(t *testing.T) {
	ioPath := NewIOPath()
	res := newIONode(ioPath, "a", "b.txt").GetParent()
	assert.Equal(
		t,
		res.GetPath(),
		"a",
	)
}

func Test_ioNode_IsExist_F(t *testing.T) {
	ioPath := NewIOPath()
	res := newIONode(ioPath, "a", "b", "c").IsExist()
	assert.False(t, res)
}

func Test_ioNode_IsExist_T(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	res := newIONode(ioPath, wd, "io-node.go").IsExist()
	assert.True(t, res)
}

func Test_ioNode_Move_File(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	node := newIONode(ioPath, wd, "move.go")
	err = ioutil.WriteFile(
		node.GetPath(),
		make([]byte, 0),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer node.Remove()

	dstNode := newIONode(ioPath, wd, "move-dst.go")
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

func Test_ioNode_Move_FileIsExist(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	node := newIONode(ioPath, wd, "move.go")

	err = ioutil.WriteFile(
		node.GetPath(),
		make([]byte, 0),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer node.Remove()

	dstNode := newIONode(ioPath, wd, "move-dst.go")

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

func Test_ioNode_Remove_Directory(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	node := newIONode(ioPath, wd, "remove")
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

func Test_ioNode_Remove_DirectoryHasChildren(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	node := newIONode(ioPath, wd, "remove")
	err = os.Mkdir(
		node.GetPath(),
		os.ModePerm,
	)
	assert.NoError(t, err)

	defer os.Remove(
		node.GetPath(),
	)

	childNodePath := ioPath.Join(
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

func Test_ioNode_Remove_File(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	node := newIONode(ioPath, wd, "remove.go")
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

func Test_ioNode_Remove_不存在(t *testing.T) {
	ioPath := NewIOPath()
	err := newIONode(ioPath, "a", "b", "c").Remove()
	assert.NoError(t, err)
}
