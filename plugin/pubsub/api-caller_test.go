package pubsub

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_apiCaller_VoidCall(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(apiCaller)
		route := "a/b/c"
		body := []int{1, 2, 3}

		mockPub := NewMockIPublisher(ctrl)
		self.pub = mockPub

		mockPub.EXPECT().Publish("a-in", requestMessage{
			API:      "c",
			Body:     `[1,2,3]`,
			Endpoint: "b",
		})

		err := self.VoidCall(route, body)
		assert.NoError(t, err)
	})

	t.Run("body is nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(apiCaller)
		route := "a/b/c"

		mockPub := NewMockIPublisher(ctrl)
		self.pub = mockPub

		mockPub.EXPECT().Publish("a-in", requestMessage{
			API:      "c",
			Body:     `{}`,
			Endpoint: "b",
		})

		err := self.VoidCall(route, nil)
		assert.NoError(t, err)
	})
}
