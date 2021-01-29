package redisex

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_redisMQ_Publish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(redisMQ)
	channel := "c"
	message := "msg"

	mockRedis := NewMockIRedis(ctrl)
	self.Redis = mockRedis

	mockRedis.EXPECT().Publish(channel, message).Return(0, nil)

	err := self.Publish(channel, message)
	assert.NoError(t, err)
}

func Test_redisMQ_Subscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(redisMQ)
	channel := "c"

	mockRedis := NewMockIRedis(ctrl)
	self.Redis = mockRedis

	mockRedis.EXPECT().Subscribe(
		[]string{channel},
		gomock.Any(),
	)

	msg := make(chan string)
	defer close(msg)
	self.Subscribe(channel, msg)
}
