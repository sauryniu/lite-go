//go:generate mockgen -destination i-sender_mock.go -package netex github.com/ahl5esoft/lite-go/netex ISender

package netex

// ISender is 发送方
type ISender interface {
	Send(route string, body interface{}) error
	SendAndReceive(route string, body interface{}) (interface{}, error)
}
