package redisex

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_sender_Send(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(sender)
		route := "a/b/c"
		body := []int{1, 2, 3}

		mockRedis := NewMockIRedis(ctrl)
		self.redis = mockRedis

		mockRedis.EXPECT().Publish("a-in", senderSendRequest{
			API:      "c",
			Body:     `[1,2,3]`,
			Endpoint: "b",
		})

		err := self.Send(route, body)
		assert.NoError(t, err)
	})

	t.Run("body is nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(sender)
		route := "a/b/c"

		mockRedis := NewMockIRedis(ctrl)
		self.redis = mockRedis

		mockRedis.EXPECT().Publish("a-in", senderSendRequest{
			API:      "c",
			Body:     `{}`,
			Endpoint: "b",
		})

		err := self.Send(route, nil)
		assert.NoError(t, err)
	})
}
