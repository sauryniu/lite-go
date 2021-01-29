package mq

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_sender_Send(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(sender)
	route := "a/b/c"
	body := []int{1, 2, 3}

	mockMQ := NewMockIMQ(ctrl)
	self.messageQueue = mockMQ

	mockMQ.EXPECT().Publish("a-in", sendRequest{
		API:      "c",
		Body:     `[1,2,3]`,
		Endpoint: "b",
	})

	err := self.Send(route, body)
	assert.NoError(t, err)
}

func Test_sender_Send_BodyIsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	self := new(sender)
	route := "a/b/c"

	mockMQ := NewMockIMQ(ctrl)
	self.messageQueue = mockMQ

	mockMQ.EXPECT().Publish("a-in", sendRequest{
		API:      "c",
		Body:     `{}`,
		Endpoint: "b",
	})

	err := self.Send(route, nil)
	assert.NoError(t, err)
}

// gomock not support chan
// func Test_sender_SendAndReceive(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	self := new(sender)
// 	route := "a/b/c"
// 	out := []int{1, 2, 3}
// 	replyID := "rid"

// 	mockIDGenertor := object.NewMockIStringGenerator(ctrl)
// 	self.idGenerator = mockIDGenertor
// 	mockMQ := NewMockIMQ(ctrl)
// 	self.messageQueue = mockMQ

// 	mockIDGenertor.EXPECT().Generate().Return(replyID)

// 	res, err := self.SendAndReceive(route, nil)
// 	assert.NoError(t, err)
// 	assert.Equal(t, res, out)
// }
