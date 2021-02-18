package osex

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ioDirectory_Create_不存在(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	dir := NewIODirectory(ioPath, cwd, "not-exist")
	err = dir.Create()
	assert.NoError(t, err)

	err = os.Remove(
		dir.GetPath(),
	)
	assert.NoError(t, err)
}

func Test_ioDirectory_FindDirectories(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	childDirPath := ioPath.Join(cwd, "dir")
	err = os.Mkdir(childDirPath, os.ModePerm)
	assert.NoError(t, err)

	defer os.Remove(childDirPath)

	res := NewIODirectory(ioPath, cwd).FindDirectories()
	assert.Len(t, res, 1)
}

func Test_ioDirectory_FindFiles(t *testing.T) {
	cwd, err := os.Getwd()
	assert.NoError(t, err)

	ioPath := NewIOPath()
	childDirPath := ioPath.Join(cwd, "files")
	err = os.Mkdir(childDirPath, os.ModePerm)
	assert.NoError(t, err)

	defer os.Remove(childDirPath)

	res := NewIODirectory(ioPath, childDirPath).FindFiles()
	assert.Len(t, res, 0)
}

func Test_ioDirectory_FindDirectories_NotExists(t *testing.T) {
	ioPath := NewIOPath()
	res := NewIODirectory(ioPath, "a", "b", "c").FindDirectories()
	assert.Len(t, res, 0)
}

func Test_ioDirectory_GetName(t *testing.T) {
	ioPath := NewIOPath()
	res := NewIODirectory(ioPath, "a", "b", "c").GetName()
	assert.Equal(t, res, "c")
}
