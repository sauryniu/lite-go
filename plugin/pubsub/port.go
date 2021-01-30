package pubsub

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

type listener struct {
	project  string
	subMsg   chan Message
	pub      IPublisher
	sub      ISubscriber
	validate *validator.Validate
}

func (m listener) Listen() {
	m.sub.Subscribe([]string{
		fmt.Sprintf(subChannelFormat, m.project),
	}, m.subMsg)
	fmt.Println(
		m.project,
		"启动于",
		time.Now().Format("2006-01-02 15:04:05"),
	)

	for {
		m.receive()
	}
}

func (m listener) receive() {
	subMsg := <-m.subMsg
	var err error
	var msg responseMessage
	if err = jsoniter.UnmarshalFromString(subMsg.Text, &msg); err != nil {
		return
	}

	var resp api.Response
	resp.Data = ""
	defer func() {
		if rv := recover(); rv != nil {
			if cErr, ok := rv.(error); ok {
				err = cErr
			} else {
				err = fmt.Errorf("%v", rv)
			}
		}

		if err != nil {
			if cErr, ok := err.(errorex.Custom); ok {
				resp.Error = cErr.Code
				resp.Data = cErr.Error()
			} else {
				fmt.Println(
					fmt.Sprintf("%v", err),
				)
				resp.Error = errorex.PanicCode
			}
		}

		if msg.ReplyID == "" {
			return
		}

		m.pub.Publish(
			fmt.Sprintf(pubChannelFormat, m.project),
			replyMessage{
				Data:    resp,
				ReplyID: msg.ReplyID,
			},
		)
	}()

	apiInstance := api.New(msg.Endpoint, msg.API)
	jsoniter.UnmarshalFromString(msg.Body, apiInstance)
	if err = m.validate.Struct(apiInstance); err != nil {
		err = errorex.New(errorex.VerifyCode, "")
		return
	}

	resp.Data, err = apiInstance.Call()
}

// NewPort is 发布订阅端口
func NewPort(project string, sub ISubscriber, pub IPublisher) api.IPort {
	return &listener{
		project:  project,
		pub:      pub,
		sub:      sub,
		validate: validator.New(),
	}
}
