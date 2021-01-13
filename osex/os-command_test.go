package osex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_osCommand_Exec(t *testing.T) {
	stdout, stderr, err := NewOSCommand("go", "version").Exec()
	assert.NoError(t, err)
	assert.Contains(
		t,
		stdout,
		"go version",
	)
	assert.Empty(t, stderr)
}

func Test_osCommand_Exec_过期(t *testing.T) {
	stdout, stderr, err := NewOSCommand("go", "version").SetExpires(0 * time.Second).Exec()
	assert.NoError(t, err)
	assert.Empty(t, stdout)
	assert.Equal(t, stderr, expiredText)
}

func Test_osCommand_Exec_No_Output(t *testing.T) {
	stdout, stderr, err := NewOSCommand("powershell", "cd").Exec()
	assert.NoError(t, err)
	assert.Empty(t, stdout)
	assert.Empty(t, stderr)
}

func Test_osCommand_Exec_Stderr(t *testing.T) {
	stdout, stderr, err := NewOSCommand("cd").Exec()
	assert.Error(t, err)
	assert.Empty(t, stdout)
	assert.Empty(t, stderr)
}
