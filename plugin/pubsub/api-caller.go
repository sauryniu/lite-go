package pubsub

import (
	"fmt"
	"strings"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/ahl5esoft/lite-go/object"
	jsoniter "github.com/json-iterator/go"
)

const emptyJSON = `{}`

type apiCaller struct {
	idGenerator object.IStringGenerator
	pub         IPublisher
	sub         ISubscriber
}

func (m apiCaller) Call(route string, body interface{}, expires time.Duration) (res interface{}, err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyString(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	req := apiMessage{
		API:      routeParams[2],
		Body:     bodyJSON,
		Endpoint: routeParams[1],
		ReplyID:  m.idGenerator.Generate(),
	}
	subChannel := fmt.Sprintf("%s-%s", routeParams[0], req.ReplyID)
	subMsg := make(chan Message)
	m.sub.Subscribe([]string{subChannel}, subMsg)
	defer m.sub.Unsubscribe(subChannel)

	m.pub.Publish(routeParams[0], req)

	var resp api.Response
	select {
	case <-time.After(expires):
		resp.Data = ""
		resp.Error = errorex.PanicCode
	case msg := <-subMsg:
		jsoniter.UnmarshalFromString(msg.Text, &resp)
	}

	if resp.Error != 0 {
		return nil, errorex.New(
			resp.Error,
			resp.Data.(string),
		)
	}

	return resp.Data, nil
}

func (m apiCaller) VoidCall(route string, body interface{}) (err error) {
	var bodyJSON string
	if bodyJSON, err = m.getBodyString(body); err != nil {
		return
	}

	routeParams := strings.Split(route, "/")
	_, err = m.pub.Publish(routeParams[0], apiMessage{
		API:      routeParams[2],
		Body:     bodyJSON,
		Endpoint: routeParams[1],
	})
	return
}

func (m apiCaller) getBodyString(body interface{}) (string, error) {
	if body == nil {
		return emptyJSON, nil
	}

	if s, ok := body.(string); ok {
		return s, nil
	}

	if bf, ok := body.([]byte); ok {
		return string(bf), nil
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
