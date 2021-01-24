package mq

// IMQ is 消息队列接口
type IMQ interface {
	Publish(channel string, message interface{}) error
	Subscribe(channel string, message chan<- string)
}
