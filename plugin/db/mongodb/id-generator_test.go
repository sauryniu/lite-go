package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewIDGenerator_Generate(t *testing.T) {
	res := NewIDGenerator().Generate()
	assert.NotEmpty(t, res)
}
