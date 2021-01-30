//go:generate mockgen -destination i-subscriber_mock.go -package pubsub github.com/ahl5esoft/lite-go/plugin/pubsub ISubscriber

package pubsub

// Message is 订阅消息
type Message struct {
	Channel string
	Text    string
}

// ISubscriber is 订阅接口
type ISubscriber interface {
	Subscribe(channels []string, message chan<- Message)
	Unsubscribe(channels ...string) error
}
