package log

import (
	"testing"

	"github.com/ahl5esoft/lite-go/contract"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_factory_Build(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockLog := contract.NewMockILog(ctrl)
		self := NewFactory(func() contract.ILog {
			return mockLog
		})

		res := self.Build()
		assert.Equal(t, res, mockLog)
	})
}
