package cor

import (
	"errors"
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/ahl5esoft/lite-go/reflectex"
	"github.com/stretchr/testify/assert"
)

type testContext struct {
	count int
}

type testHandler struct {
	IHandler

	count int
	ctx   *testContext
}

func (m *testHandler) Handle() error {
	if m.count < 0 {
		return errors.New("err")
	}

	m.ctx.count = m.ctx.count + m.count
	return m.IHandler.Handle()
}

type testBreakHandler struct {
	IHandler

	ctx *testContext
}

func (m testBreakHandler) Handle() error {
	m.Break()
	return m.IHandler.Handle()
}

type iCounter interface {
	Count() int
}

type testCounter int

func (m testCounter) Count() int {
	return int(m)
}

type testCounterHandler struct {
	IHandler

	Counter iCounter `inject:""`

	ctx *testContext
}

func (m testCounterHandler) Handle() error {
	m.ctx.count = m.Counter.Count()
	return m.IHandler.Handle()
}

func Test_Handler_Handle(t *testing.T) {
	ctx := new(testContext)
	err := (&testHandler{
		IHandler: New(),
		count:    5,
		ctx:      ctx,
	}).Handle()
	assert.NoError(t, err)
	assert.Equal(t, ctx.count, 5)
}

func Test_Handler_SetNext(t *testing.T) {
	ctx := new(testContext)
	h := &testHandler{
		IHandler: New(),
		count:    10,
		ctx:      ctx,
	}
	h.SetNext(&testHandler{
		IHandler: New(),
		count:    5,
		ctx:      ctx,
	})
	err := h.Handle()
	assert.NoError(t, err)
	assert.Equal(t, ctx.count, 15)
}

func Test_handler_SetNext_第一个错误(t *testing.T) {
	ctx := new(testContext)
	h := &testHandler{
		IHandler: New(),
		count:    -1,
		ctx:      ctx,
	}
	h.SetNext(&testHandler{
		IHandler: New(),
		count:    5,
		ctx:      ctx,
	})
	err := h.Handle()
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		"err",
	)
	assert.Equal(t, ctx.count, 0)
}

func Test_handler_SetNext_第一个跳出(t *testing.T) {
	ctx := new(testContext)
	h := &testBreakHandler{
		IHandler: New(),
		ctx:      ctx,
	}
	h.SetNext(&testHandler{
		IHandler: New(),
		count:    5,
		ctx:      ctx,
	})
	err := h.Handle()
	assert.NoError(t, err)
	assert.Equal(t, ctx.count, 0)
}

func Test_handler_SetNext_第二个错误(t *testing.T) {
	ctx := new(testContext)
	h := &testHandler{
		IHandler: New(),
		count:    5,
		ctx:      ctx,
	}
	h.SetNext(&testHandler{
		IHandler: New(),
		count:    -1,
		ctx:      ctx,
	})
	err := h.Handle()
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		"err",
	)
	assert.Equal(t, ctx.count, 5)
}

func Test_Handler_SetNext_注入(t *testing.T) {
	counterType := reflectex.InterfaceTypeOf(
		(*iCounter)(nil),
	)
	ioc.Set(
		counterType,
		testCounter(11),
	)

	ctx := new(testContext)
	h := &testHandler{
		IHandler: New(),
		count:    0,
		ctx:      ctx,
	}
	h.SetNext(&testCounterHandler{
		IHandler: New(),
		ctx:      ctx,
	})

	err := h.Handle()
	assert.NoError(t, err)
	assert.Equal(t, ctx.count, 11)
}
