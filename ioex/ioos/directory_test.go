package ioos

import (
	"os"
	"testing"

	"github.com/ahl5esoft/lite-go/ioex/iopath"
	"github.com/stretchr/testify/assert"
)

func Test_directory_Create_不存在(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	dir := NewDirectory(cwd, "not-exist")
	err = dir.Create()
	assert.NoError(t, err)

	err = os.Remove(
		dir.GetPath(),
	)
	assert.NoError(t, err)
}

func Test_directory_FindDirectories(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	childDirPath := iopath.Join(cwd, "dir")
	err = os.Mkdir(childDirPath, os.ModePerm)
	assert.NoError(t, err)

	defer os.Remove(childDirPath)

	res := NewDirectory(cwd).FindDirectories()
	assert.Len(t, res, 1)
}

func Test_directory_FindFiles(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	childDirPath := iopath.Join(cwd, "files")
	err = os.Mkdir(childDirPath, os.ModePerm)
	assert.NoError(t, err)

	defer os.Remove(childDirPath)

	res := NewDirectory(childDirPath).FindFiles()
	assert.Len(t, res, 0)
}

func Test_directory_FindDirectories_NotExists(t *testing.T) {
	res := NewDirectory("a", "b", "c").FindDirectories()
	assert.Len(t, res, 0)
}

func Test_directory_GetName(t *testing.T) {
	res := NewDirectory("a", "b", "c").GetName()
	assert.Equal(t, res, "c")
}
