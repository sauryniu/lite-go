package randex

import (
	"testing"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/stretchr/testify/assert"
)

func Test_NewStringGenerator(t *testing.T) {
	length := 10
	res := NewStringGenerator(length).Generate()
	assert.Len(t, res, length)

	ok := underscore.Chain(
		[]byte(res),
	).All(func(r byte, _ int) bool {
		return underscore.Chain(defaultStringGeneratorSource).Any(func(cr byte, _ int) bool {
			return r == cr
		})
	})
	assert.True(t, ok)
}

func Test_NewStringGenerator_NewStringGeneratorSourceOption(t *testing.T) {
	length := 10
	source := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	res := NewStringGenerator(
		length,
		NewStringGeneratorSourceOption(source),
	).Generate()
	assert.Len(t, res, length)

	ok := underscore.Chain(
		[]byte(res),
	).All(func(r byte, _ int) bool {
		return underscore.Chain(source).Any(func(cr byte, _ int) bool {
			return r == cr
		})
	})
	assert.True(t, ok)
}
