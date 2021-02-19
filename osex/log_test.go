package osex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_osLog_Debug(t *testing.T) {
	NewLog().AddAttr("p1", "%s-%s", "a", "b").Debug()
	assert.True(t, false)
}
