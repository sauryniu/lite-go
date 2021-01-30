package redisex

import (
	"fmt"
	"strings"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/ahl5esoft/lite-go/netex"
	"github.com/ahl5esoft/lite-go/object"
	"github.com/ahl5esoft/lite-go/plugin/pubsub"
	jsoniter "github.com/json-iterator/go"
)

const emptyJSON = `{}`

type senderSendResponse struct {
	Data    api.Response
	ReplyID string
}

type senderSendRequest struct {
	API      string
	Body     string
	Endpoint string
	ReplyID  string
}

type sender struct {
	idGenerator object.IStringGenerator
	redis       IRedis
}

func (m sender) Send(route string, body interface{}) (err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyJSON(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	_, err = m.redis.Publish(
		fmt.Sprintf("%s-in", routeParams[0]),
		senderSendRequest{
			API:      routeParams[2],
			Body:     bodyJSON,
			Endpoint: routeParams[1],
		},
	)
	return
}

func (m sender) SendAndReceive(route string, body interface{}) (res interface{}, err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyJSON(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	outChannel := fmt.Sprintf("%s-out", routeParams[0])
	subMsg := make(chan pubsub.Message)
	m.redis.Subscribe([]string{outChannel}, subMsg)

	req := senderSendRequest{
		API:      routeParams[2],
		Body:     bodyJSON,
		Endpoint: routeParams[1],
		ReplyID:  m.idGenerator.Generate(),
	}
	fmt.Println(req)
	m.redis.Publish(
		fmt.Sprintf("%s-in", routeParams[0]),
		req,
	)

	var resp senderSendResponse
	for {
		jsoniter.UnmarshalFromString(
			(<-subMsg).Text,
			&resp,
		)
		fmt.Println(resp)
		if resp.ReplyID == req.ReplyID {
			m.redis.Unsubscribe(outChannel)
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

// NewSender is redis发送者
func NewSender(redis IRedis, idGenerator object.IStringGenerator) netex.ISender {
	return &sender{
		idGenerator: idGenerator,
		redis:       redis,
	}
}
