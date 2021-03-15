package pubsub

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/ahl5esoft/lite-go/log"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

type apiPort struct {
	apiFactory api.IFactory
	log        log.ILog
	project    string
	pub        IPublisher
	sub        ISubscriber
	subMsg     chan Message
	validate   *validator.Validate
}

func (m apiPort) Listen() {
	m.sub.Subscribe([]string{m.project}, m.subMsg)
	m.log.AddDesc(m.project).AddAttr(
		"run-at",
		time.Now().Format("2006-01-02 03:04:05"),
	).Info()

	for {
		m.handle()
	}
}

func (m apiPort) handle() {
	subMsg := <-m.subMsg

	var err error
	var req apiMessage
	if err = jsoniter.UnmarshalFromString(subMsg.Text, &req); err != nil {
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
				resp.Error = errorex.PanicCode
			}
		}

		if req.ReplyID == "" {
			return
		}

		m.pub.Publish(
			fmt.Sprintf("%s-%s", m.project, req.ReplyID),
			resp,
		)
	}()

	apiInstance := m.apiFactory.Build(req.Endpoint, req.API)
	jsoniter.UnmarshalFromString(req.Body, apiInstance)
	if err = m.validate.Struct(apiInstance); err != nil {
		m.log.AddAttr("validate", "%v", err).Error()
		err = errorex.New(errorex.VerifyCode, "")
		return
	}

	resp.Data, err = apiInstance.Call()
}

// NewAPIPort is 发布订阅端口
func NewAPIPort(project string, sub ISubscriber, pub IPublisher, apiFactory api.IFactory, log log.ILog) api.IPort {
	return &apiPort{
		apiFactory: apiFactory,
		log:        log,
		project:    project,
		pub:        pub,
		sub:        sub,
		subMsg:     make(chan Message),
		validate:   validator.New(),
	}
}
