package pubsub

type apiMessage struct {
	API      string `json:"api"`
	Body     string `json:"body"`
	Endpoint string `json:"endpoint"`
	ReplyID  string `json:"replyID"`
}
