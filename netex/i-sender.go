package netex

// ISender is 发送方
type ISender interface {
	Send(route string, body interface{}) error
	SendAndReceive(route string, body interface{}) (interface{}, error)
}
