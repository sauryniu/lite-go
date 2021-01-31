package pubsub

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

const (
	apiPortPubChannelFormat = "%s-out"
	apiPortSubChannelFormat = "%s-in"
)

type apiPort struct {
	apiFactory api.IFactory
	project    string
	pub        IPublisher
	sub        ISubscriber
	subMsg     chan Message
	validate   *validator.Validate
}

func (m apiPort) Listen() {
	m.sub.Subscribe([]string{
		fmt.Sprintf(apiPortSubChannelFormat, m.project),
	}, m.subMsg)
	fmt.Println(
		m.project,
		"启动于",
		time.Now().Format("2006-01-02 15:04:05"),
	)

	for {
		m.handle()
	}
}

func (m apiPort) handle() {
	subMsg := <-m.subMsg
	var err error
	var msg requestMessage
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
			fmt.Sprintf(apiPortPubChannelFormat, m.project),
			responseMessage{
				Data:    resp,
				ReplyID: msg.ReplyID,
			},
		)
	}()

	apiInstance := m.apiFactory.Build(msg.Endpoint, msg.API)
	jsoniter.UnmarshalFromString(msg.Body, apiInstance)
	if err = m.validate.Struct(apiInstance); err != nil {
		err = errorex.New(errorex.VerifyCode, "")
		return
	}

	resp.Data, err = apiInstance.Call()
}

// NewPort is 发布订阅端口
func NewPort(project string, sub ISubscriber, pub IPublisher, apiFactory api.IFactory) api.IPort {
	return &apiPort{
		apiFactory: apiFactory,
		project:    project,
		pub:        pub,
		sub:        sub,
		validate:   validator.New(),
	}
}
