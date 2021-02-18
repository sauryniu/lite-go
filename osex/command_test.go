package osex

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_osCommand_Exec(t *testing.T) {
	stdout, stderr, err := NewCommand().Exec("go", "version")
	assert.NoError(t, err)
	assert.Contains(
		t,
		stdout,
		"go version",
	)
	assert.Empty(t, stderr)
}

func Test_osCommand_Exec_SetDir(t *testing.T) {
	wd, _ := os.Getwd()
	ioPath := NewIOPath()
	dir := NewIODirectory(ioPath, wd, "set-dir")
	dir.Create()
	defer dir.Remove()

	_, _, err := NewCommand().SetDir(
		dir.GetPath(),
	).Exec("git", "init")
	assert.NoError(t, err)

	ok := NewIODirectory(
		ioPath,
		dir.GetPath(),
		".git",
	).IsExist()
	assert.True(t, ok)
}

func Test_osCommand_Exec_过期(t *testing.T) {
	stdout, stderr, err := NewCommand().SetExpires(0*time.Second).Exec("go", "version")
	assert.NoError(t, err)
	assert.Empty(t, stdout)
	assert.Equal(t, stderr, expiredText)
}

func Test_osCommand_Exec_No_Output(t *testing.T) {
	stdout, stderr, err := NewCommand().Exec("powershell", "cd")
	assert.NoError(t, err)
	assert.Empty(t, stdout)
	assert.Empty(t, stderr)
}

func Test_osCommand_Exec_Stderr(t *testing.T) {
	stdout, stderr, err := NewCommand().Exec("cd")
	assert.Error(t, err)
	assert.Empty(t, stdout)
	assert.Empty(t, stderr)
}

func Test_osCommand_SetDir(t *testing.T) {
	cmd := NewCommand().SetDir("dir-path")
	assert.Equal(
		t,
		cmd.(*osCommand).dirPath,
		"dir-path",
	)
}

func Test_osCommand_SetExpires(t *testing.T) {
	cmd := NewCommand().SetExpires(time.Second)
	assert.Equal(
		t,
		cmd.(*osCommand).expires,
		time.Second,
	)
}
