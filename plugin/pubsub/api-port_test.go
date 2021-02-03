package pubsub

import (
	"fmt"
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

type testAPI struct{}

func (m testAPI) Call() (interface{}, error) {
	return "ok", nil
}

func Test_apiPort_handle(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		self := new(apiPort)
		self.project = "project"
		self.subMsg = make(chan Message)
		self.validate = validator.New()

		mockAPIFactory := api.NewMockIFactory(ctrl)
		self.apiFactory = mockAPIFactory
		mockPub := NewMockIPublisher(ctrl)
		self.pub = mockPub

		mockAPIFactory.EXPECT().Build("endpoint", "api").Return(
			new(testAPI),
		)

		replyID := "r-id"
		mockPub.EXPECT().Publish(
			fmt.Sprintf("%s-%s", self.project, replyID),
			api.Response{
				Data:  "ok",
				Error: 0,
			},
		)

		go func() {
			self.subMsg <- Message{
				Text: fmt.Sprintf(`{"API":"api","Body":"{}","Endpoint":"endpoint","ReplyID":"%s"}`, replyID),
			}
		}()

		self.handle()
	})

	t.Run("Message.Text is invalid json", func(t *testing.T) {
		self := new(apiPort)
		self.subMsg = make(chan Message)

		go func() {
			self.subMsg <- Message{
				Text: "",
			}
		}()

		self.handle()
	})
}
