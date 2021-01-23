package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewIDGenerator_Generate(t *testing.T) {
	res := NewStringGenerator().Generate()
	assert.NotEmpty(t, res)
	assert.Len(t, res, 24)
}
