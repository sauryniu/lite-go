package cor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testContext struct {
	Count int

	isBreak bool
}

func (m *testContext) Break() {
	m.isBreak = true
}

func (m testContext) IsBreak() bool {
	return m.isBreak
}

func Test_handler_Handle(t *testing.T) {
	ctx := new(testContext)
	err := New(func(ctx interface{}) error {
		ctx.(*testContext).Count = 5
		return nil
	}).Handle(ctx)
	assert.NoError(t, err)
	assert.Equal(t, ctx.Count, 5)
}

func Test_handler_SetNext(t *testing.T) {
	ctx := new(testContext)
	h := New(func(ctx interface{}) error {
		ctx.(*testContext).Count = 5
		return nil
	})
	h.SetNext(
		New(func(ctx interface{}) error {
			ctx.(*testContext).Count = ctx.(*testContext).Count + 10
			return nil
		}),
	)
	err := h.Handle(ctx)
	assert.NoError(t, err)
	assert.Equal(t, ctx.Count, 15)
}

func Test_handler_SetNext_第一个错误(t *testing.T) {
	ctx := new(testContext)
	h := New(func(ctx interface{}) error {
		return errors.New("err")
	})
	h.SetNext(
		New(func(ctx interface{}) error {
			ctx.(*testContext).Count = 5
			return nil
		}),
	)
	err := h.Handle(ctx)
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		"err",
	)
	assert.Equal(t, ctx.Count, 0)
}

func Test_handler_SetNext_第一个跳出(t *testing.T) {
	ctx := new(testContext)
	h := New(func(ctx interface{}) error {
		ctx.(*testContext).Break()
		return nil
	})
	h.SetNext(
		New(func(ctx interface{}) error {
			ctx.(*testContext).Count = 5
			return nil
		}),
	)
	err := h.Handle(ctx)
	assert.NoError(t, err)
	assert.Equal(t, ctx.Count, 0)
}

func Test_handler_SetNext_第二个错误(t *testing.T) {
	ctx := new(testContext)
	h := New(func(ctx interface{}) error {
		ctx.(*testContext).Count = 5
		return nil
	})
	h.SetNext(
		New(func(ctx interface{}) error {
			return errors.New("err")
		}),
	)
	err := h.Handle(ctx)
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		"err",
	)
	assert.Equal(t, ctx.Count, 5)
}
