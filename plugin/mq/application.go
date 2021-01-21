package mq

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/ahl5esoft/lite-go/runtimeex"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

type receiveMessage struct {
	API      string
	Body     string
	Endpoint string
	ReplyID  string
}

type replyMessage struct {
	Data    api.Response
	ReplyID string
}

type application struct {
	mq         IMQ
	outChannel string
	project    string
	validate   *validator.Validate
}

func (m application) Run() {
	msg := make(chan string)
	m.mq.Subscribe(
		fmt.Sprintf("%s-in", m.project),
		msg,
	)
	fmt.Println(
		fmt.Sprintf(
			"启动于%s",
			time.Now().Format("2006-01-02 15:04:05"),
		),
	)

	for {
		m.receive(<-msg)
	}
}

func (m application) getOutChannel() string {
	if m.outChannel == "" {
		m.outChannel = fmt.Sprintf("%s-out", m.project)
	}

	return m.outChannel
}

func (m application) receive(message string) {
	var err error
	var rm receiveMessage
	if err = jsoniter.UnmarshalFromString(message, &rm); err != nil {
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

		m.mq.Publish(
			m.getOutChannel(),
			replyMessage{
				Data:    resp,
				ReplyID: rm.ReplyID,
			},
		)
	}()

	apiInstance := api.New(rm.Endpoint, rm.API)
	jsoniter.UnmarshalFromString(rm.Body, apiInstance)
	if err = m.validate.Struct(apiInstance); err != nil {
		err = errorex.New(errorex.VerifyCode, "")
		return
	}

	resp.Data, err = apiInstance.Call()
}

// NewApplication is 实例化mq应用
func NewApplication(mq IMQ, project string) runtimeex.IApplication {
	return &application{
		mq:       mq,
		project:  project,
		validate: validator.New(),
	}
}
