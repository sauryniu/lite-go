package goredis

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func Test_redisMQ_Publish(t *testing.T) {
	channel := "Test_redisMQ_Publish"
	message := "hello"
	sub := client.Subscribe(channel)

	msg := make(chan (*redis.Message))
	go func() {
		for {
			select {
			case msg <- <-sub.Channel():
				sub.Close()
			default:
			}
		}
	}()

	err := NewMQ(self).Publish(channel, message)
	assert.NoError(t, err)
	assert.Equal(
		t,
		(<-msg).Payload,
		message,
	)
}

func Test_redisMQ_Subscribe(t *testing.T) {
	channel := "Test_redisMQ_Subscribe"
	message := "hello"

	res := make(chan string)
	NewMQ(self).Subscribe(channel, res)

	self.Publish(channel, message)

	assert.Equal(
		t,
		<-res,
		message,
	)
}
