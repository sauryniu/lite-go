package api

// ReceiveMessage is 接收消息
type ReceiveMessage struct {
	API      string
	Body     string
	Endpoint string
	ReplyID  string
}

// ReplyMessage is 回复消息
type ReplyMessage struct {
	Data    Response
	ReplyID string
}
