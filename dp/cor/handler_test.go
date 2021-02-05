package cor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
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

	if m.count == 0 {
		m.Break()
	}

	m.ctx.count = m.ctx.count + m.count
	return m.IHandler.Handle()
}

type iCounter interface {
	Count() int
}

type testCounter int

func (m testCounter) Count() int {
	return int(m)
}

type testInjectHandler struct {
	IHandler

	Counter iCounter `inject:""`

	ctx *testContext
}

func (m testInjectHandler) Handle() error {
	m.ctx.count = m.Counter.Count()
	return m.IHandler.Handle()
}

func Test_handler_Break(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		ctx := new(testContext)
		h := &testHandler{
			IHandler: New(),
			ctx:      ctx,
		}
		h.SetNext(&testHandler{
			IHandler: New(),
			count:    2,
			ctx:      ctx,
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 0)
	})
	t.Run("second", func(t *testing.T) {
		ctx := new(testContext)
		h := &testHandler{
			IHandler: New(),
			count:    1,
			ctx:      ctx,
		}
		h.SetNext(&testHandler{
			IHandler: New(),
			ctx:      ctx,
		}).SetNext(&testHandler{
			IHandler: New(),
			count:    3,
			ctx:      ctx,
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 1)
	})
	t.Run("ok", func(t *testing.T) {
		ctx := new(testContext)
		h := &testHandler{
			IHandler: New(),
			count:    1,
			ctx:      ctx,
		}
		h.SetNext(&testHandler{
			IHandler: New(),
			count:    2,
			ctx:      ctx,
		}).SetNext(&testHandler{
			IHandler: New(),
			count:    3,
			ctx:      ctx,
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 6)
	})
	t.Run("inject", func(t *testing.T) {
		ioc.Set(
			(*iCounter)(nil),
			testCounter(11),
		)

		ctx := new(testContext)
		h := &testHandler{
			IHandler: New(),
			count:    1,
			ctx:      ctx,
		}
		h.SetNext(&testInjectHandler{
			IHandler: New(),
			ctx:      ctx,
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 11)
	})
}

type testDelayHandler struct {
	IHandler

	count int
	ctx   *testContext
}

func (m testDelayHandler) Handle() error {
	if m.count == 0 {
		m.ctx.count = 1
	} else if m.count == 6 {
		m.ctx.count = 0
	} else {
		m.ctx.count = 6
	}
	return m.IHandler.Handle()
}

func Test_SetDelayNext(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := new(testContext)
		h := &testDelayHandler{
			IHandler: New(),
			count:    ctx.count,
			ctx:      ctx,
		}
		h.SetNext(&testDelayHandler{
			IHandler: New(),
			count:    ctx.count,
			ctx:      ctx,
		}).SetDelayNext(func() IHandler {
			fmt.Println(ctx)
			return &testDelayHandler{
				IHandler: New(),
				count:    ctx.count,
				ctx:      ctx,
			}
		})
		err := h.Handle()
		assert.NoError(t, err)
		assert.Equal(t, ctx.count, 6)
	})
}
