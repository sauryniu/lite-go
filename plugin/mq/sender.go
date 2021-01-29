package mq

import (
	"fmt"
	"strings"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/ahl5esoft/lite-go/netex"
	"github.com/ahl5esoft/lite-go/object"
	jsoniter "github.com/json-iterator/go"
)

const emptyJSON = `{}`

type sendResponse struct {
	Data    api.Response
	ReplyID string
}

type sendRequest struct {
	API      string
	Body     string
	Endpoint string
	ReplyID  string
}

type sender struct {
	idGenerator  object.IStringGenerator
	messageQueue IMQ
}

func (m sender) Send(route string, body interface{}) (err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyJSON(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	return m.messageQueue.Publish(
		fmt.Sprintf("%s-in", routeParams[0]),
		sendRequest{
			API:      routeParams[2],
			Body:     bodyJSON,
			Endpoint: routeParams[1],
		},
	)
}

func (m sender) SendAndReceive(route string, body interface{}) (res interface{}, err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyJSON(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	req := sendRequest{
		API:      routeParams[2],
		Body:     bodyJSON,
		Endpoint: routeParams[1],
		ReplyID:  m.idGenerator.Generate(),
	}

	respMessage := make(chan string)
	m.messageQueue.Subscribe(
		fmt.Sprintf("%s-out", routeParams[0]),
		respMessage,
	)

	go m.messageQueue.Publish(
		fmt.Sprintf("%s-in", routeParams[0]),
		req,
	)

	var resp sendResponse
	for {
		jsoniter.UnmarshalFromString(<-respMessage, &resp)
		if resp.ReplyID == req.ReplyID {
			break
		}
	}

	if resp.Data.Error != 0 {
		return nil, errorex.New(
			resp.Data.Error,
			resp.Data.Data.(string),
		)
	}

	return resp.Data.Data, nil
}

func (m sender) getBodyJSON(body interface{}) (string, error) {
	if body == nil {
		return emptyJSON, nil
	}

	return jsoniter.MarshalToString(body)
}

// NewSender is 消息队列发送者
func NewSender(messageQueue IMQ, idGenerator object.IStringGenerator) netex.ISender {
	return &sender{
		idGenerator:  idGenerator,
		messageQueue: messageQueue,
	}
}
