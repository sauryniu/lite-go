package pubsub

import (
	"fmt"
	"strings"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/ahl5esoft/lite-go/object"
	jsoniter "github.com/json-iterator/go"
)

const (
	emptyJSON                 = `{}`
	apiCallerPubChannelFormat = "%s-in"
	apiCallerSubChannelFormat = "%s-out"
)

type apiCaller struct {
	idGenerator object.IStringGenerator
	pub         IPublisher
	sub         ISubscriber
}

func (m apiCaller) Call(route string, body interface{}) (res interface{}, err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyJSON(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	subChannel := fmt.Sprintf(apiCallerSubChannelFormat, routeParams[0])
	subMsg := make(chan Message)
	m.sub.Subscribe([]string{subChannel}, subMsg)

	req := requestMessage{
		API:      routeParams[2],
		Body:     bodyJSON,
		Endpoint: routeParams[1],
		ReplyID:  m.idGenerator.Generate(),
	}
	m.pub.Publish(
		fmt.Sprintf(apiCallerPubChannelFormat, routeParams[0]),
		req,
	)

	var resp responseMessage
	for {
		jsoniter.UnmarshalFromString(
			(<-subMsg).Text,
			&resp,
		)
		if resp.ReplyID == req.ReplyID {
			m.sub.Unsubscribe(subChannel)
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

func (m apiCaller) VoidCall(route string, body interface{}) (err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyJSON(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	_, err = m.pub.Publish(
		fmt.Sprintf(apiCallerPubChannelFormat, routeParams[0]),
		requestMessage{
			API:      routeParams[2],
			Body:     bodyJSON,
			Endpoint: routeParams[1],
		},
	)
	return
}

func (m apiCaller) getBodyJSON(body interface{}) (string, error) {
	if body == nil {
		return emptyJSON, nil
	}

	return jsoniter.MarshalToString(body)
}

// NewAPICaller is redis api 调用
func NewAPICaller(pub IPublisher, sub ISubscriber, idGenerator object.IStringGenerator) api.ICaller {
	return &apiCaller{
		idGenerator: idGenerator,
		pub:         pub,
		sub:         sub,
	}
}
