package mq

// IoCKey is 依赖注入键
const IoCKey = "mq"

// IMQ is 消息队列接口
type IMQ interface {
	Publish(channel string, message interface{}) error
	Subscribe(channel string, message chan<- string)
}
