package pubsub

type apiMessage struct {
	API      string
	Body     string
	Endpoint string
	ReplyID  string
}
