package pubsub

import "github.com/ahl5esoft/lite-go/api"

type requestMessage struct {
	API      string
	Body     string
	Endpoint string
	ReplyID  string
}

type responseMessage struct {
	Data    api.Response
	ReplyID string
}
