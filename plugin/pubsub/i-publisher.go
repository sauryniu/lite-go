//go:generate mockgen -destination i-publisher_mock.go -package pubsub github.com/ahl5esoft/lite-go/plugin/pubsub IPublisher

package pubsub

// IPublisher is 发布接口
type IPublisher interface {
	Publish(channel string, message interface{}) (int, error)
}
