package goredis

import (
	"testing"
	"time"

	"github.com/ahl5esoft/lite-go/plugin/pubsub"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	self = &goRedis{
		client: redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}),
		subs: make(map[string]*redis.PubSub),
	}
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
)

func Test_goRedis_Del(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		key := "Test_goRedis_Del"
		client.Set(key, "test", 0).Result()

		defer client.Del(key)

		count, err := self.Del(key)
		assert.NoError(t, err)
		assert.Equal(t, count, 1)
	})

	t.Run("不存在", func(t *testing.T) {
		key := "Test_goRedis_Del_不存在"
		count, err := self.Del(key)
		assert.NoError(t, err)
		assert.Equal(t, count, 0)
	})
}

func Test_goRedis_Exists_不存在(t *testing.T) {
	t.Run("不存在", func(t *testing.T) {
		key := "Test_goRedis_Exists_不存在"
		ok, err := self.Exists(key)
		assert.NoError(t, err)
		assert.False(t, ok)
	})

	t.Run("已存在", func(t *testing.T) {
		key := "Test_goRedis_Exists_已存在"
		client.Set(key, "test", 0)
		defer client.Del(key)

		ok, err := self.Exists(key)
		assert.NoError(t, err)
		assert.True(t, ok)
	})
}

func Test_goRedis_Eval(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		key := "Test_goRedis_Eval"
		defer client.Del(key)

		value := "v"
		_, err := self.Eval(`redis.call("set", KEYS[1], ARGV[1])`, []string{key}, value)
		assert.NoError(t, err)

		time.Sleep(1 * time.Second)

		res, err := client.Get(key).Result()
		assert.NoError(t, err)
		assert.Equal(t, res, value)
	})

	t.Run("KeysIsNil", func(t *testing.T) {
		key := "Test_goRedis_Eval"
		defer client.Del(key)

		_, err := self.Eval(`redis.call("time")`, nil)
		assert.NoError(t, err)
	})
}

func Test_goRedis_Get(t *testing.T) {
	key := "Test_goRedis_Get"
	client.Set(key, "test", 0)
	res, err := self.Get(key)

	client.Del(key)

	assert.NoError(t, err)
	assert.Equal(t, res, "test")
}

func Test_goRedis_Get_不存在(t *testing.T) {
	key := "Test_goRedis_Get_不存在"
	res, err := self.Get(key)
	assert.NoError(t, err)
	assert.Empty(t, res)
}

func Test_goRedis_Publish(t *testing.T) {
	channel := "Test_goRedis_Publish"
	t.Run("string", func(t *testing.T) {
		message := "hello"
		sub := client.Subscribe(channel)

		msg := make(chan *redis.Message)
		go func() {
			for {
				msg <- <-sub.Channel()
				close(msg)
				sub.Unsubscribe(channel)
			}
		}()

		_, err := self.Publish(channel, message)
		assert.NoError(t, err)

		assert.Equal(
			t,
			(<-msg).Payload,
			message,
		)
	})

	t.Run("array", func(t *testing.T) {
		message := []int{1, 2, 3}
		sub := client.Subscribe(channel)

		msg := make(chan *redis.Message)
		go func() {
			for {
				msg <- <-sub.Channel()
				close(msg)
				sub.Unsubscribe(channel)
			}
		}()

		_, err := self.Publish(channel, message)
		assert.NoError(t, err)

		assert.Equal(
			t,
			(<-msg).Payload,
			"[1,2,3]",
		)
	})
}

func Test_goRedis_Set(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		key := "Test_goRedis_Set"
		defer client.Del(key)

		ok, err := self.Set(key, "test")
		assert.NoError(t, err)
		assert.True(t, ok)

		res, err := client.Get(key).Result()
		assert.NoError(t, err)
		assert.Equal(t, res, "test")
	})

	t.Run("NX", func(t *testing.T) {
		key := "Test_goRedis_SetNX"
		defer client.Del(key)

		ok, err := self.Set(
			key,
			"",
			"nx",
		)
		assert.NoError(t, err)
		assert.True(t, ok)

		ok, err = self.Set(
			key,
			"",
			"nx",
		)
		assert.NoError(t, err)
		assert.False(t, ok)
	})

	t.Run("XX", func(t *testing.T) {
		key := "Test_goRedis_Set_XX"
		defer client.Del(key)

		ok, err := self.Set(
			key,
			"",
			"xx",
		)
		assert.NoError(t, err)
		assert.False(t, ok)

		client.Set(key, "", 0)

		ok, err = self.Set(
			key,
			"a",
			"xx",
		)
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("EX", func(t *testing.T) {
		key := "Test_goRedis_Set_EX"
		defer self.Del(key)

		duration := 1 * time.Second
		ok, err := self.Set(key, "a", "ex", duration)
		assert.NoError(t, err)
		assert.True(t, ok)

		time.Sleep(duration * 2)

		res, err := client.Get(key).Result()
		assert.Equal(t, err, redis.Nil)
		assert.Equal(t, res, "")
	})

	t.Run("EX_NX", func(t *testing.T) {
		key := "Test_goRedis_Set_EX_NX"
		defer self.Del(key)

		duration := 1 * time.Second
		ok, err := self.Set(key, "a", "ex", duration, "nx")
		assert.NoError(t, err)
		assert.True(t, ok)

		res, err := client.Get(key).Result()
		assert.NoError(t, err)
		assert.Equal(t, res, "a")

		ok, err = self.Set(key, "a", "ex", duration, "nx")
		assert.NoError(t, err)
		assert.False(t, ok)

		time.Sleep(duration * 2)

		ok, err = self.Set(key, "a", "ex", duration, "nx")
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("EX_XX", func(t *testing.T) {
		key := "Test_goRedis_Set_EX_XX"
		defer self.Del(key)

		duration := 1 * time.Second
		ok, err := self.Set(key, "a", "ex", duration, "xx")
		assert.NoError(t, err)
		assert.False(t, ok)

		client.Set(key, "", 0)

		ok, err = self.Set(key, "a", "ex", duration, "xx")
		assert.NoError(t, err)
		assert.True(t, ok)

		time.Sleep(duration * 2)

		_, err = client.Get(key).Result()
		assert.Equal(t, err, redis.Nil)
	})

	t.Run("PX", func(t *testing.T) {
		key := "Test_goRedis_SetNX_PX"
		defer self.Del(key)

		duration := 50 * time.Millisecond
		ok, err := self.Set(key, "a", "px", duration)
		assert.NoError(t, err)
		assert.True(t, ok)

		time.Sleep(duration * 2)

		res, err := client.Get(key).Result()
		assert.Equal(t, err, redis.Nil)
		assert.Equal(t, res, "")
	})

	t.Run("PX_NX", func(t *testing.T) {
		key := "Test_goRedis_Set_PX_NX"
		defer self.Del(key)

		duration := 300 * time.Millisecond
		ok, err := self.Set(key, "a", "px", duration, "nx")
		assert.NoError(t, err)
		assert.True(t, ok)

		res, err := client.Get(key).Result()
		assert.NoError(t, err)
		assert.Equal(t, res, "a")

		ok, err = self.Set(key, "a", "px", duration, "nx")
		assert.NoError(t, err)
		assert.False(t, ok)

		time.Sleep(duration * 2)

		ok, err = self.Set(key, "a", "px", duration, "nx")
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("PX_XX", func(t *testing.T) {
		key := "Test_goRedis_Set_PX_XX"
		defer self.Del(key)

		duration := 300 * time.Millisecond
		ok, err := self.Set(key, "a", "px", duration, "xx")
		assert.NoError(t, err)
		assert.False(t, ok)

		client.Set(key, "", 0)

		ok, err = self.Set(key, "a", "px", duration, "xx")
		assert.NoError(t, err)
		assert.True(t, ok)

		time.Sleep(duration * 2)

		_, err = client.Get(key).Result()
		assert.Equal(t, err, redis.Nil)
	})
}

func Test_goRedis_Subscribe(t *testing.T) {
	channel := "Test_goRedis_Subscribe"
	msg := make(chan pubsub.Message)
	defer close(msg)
	self.Subscribe([]string{channel}, msg)

	message := "hello"
	count, err := client.Publish(channel, message).Result()
	assert.NoError(t, err)
	assert.Equal(
		t,
		count,
		int64(1),
	)
	assert.Equal(
		t,
		(<-msg).Text,
		message,
	)
}

func Test_goRedis_Unsubscribe(t *testing.T) {
	channel := "Test_goRedis_Unsubscribe"

	res := make(chan string)
	sub := client.Subscribe(channel)
	go func() {
		msg := <-sub.Channel()
		res <- msg.Payload
		close(res)
	}()

	self.subs = map[string]*redis.PubSub{
		channel: sub,
	}
	err := self.Unsubscribe(channel)
	assert.NoError(t, err)
	assert.Len(t, self.subs, 0)

	message := "hello"
	count, err := client.Publish(channel, message).Result()
	assert.NoError(t, err)
	assert.Equal(
		t,
		count,
		int64(0),
	)
}
