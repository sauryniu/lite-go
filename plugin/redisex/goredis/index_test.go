package goredis

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	self   = New("127.0.0.1", 6379)
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
)

func Test_goRedis_Del(t *testing.T) {
	key := "Test_goRedis_Del"
	client.Set(key, "test", 0).Result()

	defer client.Del(key)

	count, err := self.Del(key)
	assert.NoError(t, err)
	assert.Equal(t, count, 1)
}

func Test_goRedis_Del_不存在(t *testing.T) {
	key := "Test_goRedis_Del_不存在"
	count, err := self.Del(key)
	assert.NoError(t, err)
	assert.Equal(t, count, 0)
}

func Test_goRedis_Exists_不存在(t *testing.T) {
	key := "Test_goRedis_Exists_不存在"
	ok, err := self.Exists(key)
	assert.NoError(t, err)
	assert.False(t, ok)
}

func Test_goRedis_Exists_已存在(t *testing.T) {
	key := "Test_goRedis_Exists_已存在"
	client.Set(key, "test", 0)
	defer client.Del(key)

	ok, err := self.Exists(key)
	assert.NoError(t, err)
	assert.True(t, ok)
}

func Test_goRedis_Eval(t *testing.T) {
	key := "Test_goRedis_Eval"
	defer client.Del(key)

	value := "v"
	_, err := self.Eval(`redis.call("set", KEYS[1], ARGV[1])`, []string{key}, value)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	res, err := client.Get(key).Result()
	assert.NoError(t, err)
	assert.Equal(t, res, value)
}

func Test_goRedis_Eval_KeysIsNil(t *testing.T) {
	key := "Test_goRedis_Eval"
	defer client.Del(key)

	_, err := self.Eval(`redis.call("time")`, nil)
	assert.NoError(t, err)
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

func Test_goRedis_Publish_string(t *testing.T) {
	channel := "Test_goRedis_Publish"
	message := "hello"
	sub := client.Subscribe(channel)

	msg := make(chan *redis.Message)
	go func() {
		for {
			select {
			case msg <- <-sub.Channel():
				sub.Close()
			default:
			}
		}
	}()

	_, err := self.Publish(channel, message)
	assert.NoError(t, err)

	assert.Equal(
		t,
		(<-msg).Payload,
		message,
	)
}

func Test_goRedis_Publish_array(t *testing.T) {
	channel := "Test_goRedis_Publish"
	message := []int{1, 2, 3}
	sub := client.Subscribe(channel)

	msg := make(chan *redis.Message)
	go func() {
		for {
			select {
			case msg <- <-sub.Channel():
				sub.Close()
			default:
			}
		}
	}()

	_, err := self.Publish(channel, message)
	assert.NoError(t, err)

	assert.Equal(
		t,
		(<-msg).Payload,
		"[1,2,3]",
	)
}

func Test_goRedis_Set(t *testing.T) {
	key := "Test_goRedis_Set"
	defer client.Del(key)

	ok, err := self.Set(key, "test")
	assert.NoError(t, err)
	assert.True(t, ok)

	res, err := client.Get(key).Result()
	assert.NoError(t, err)
	assert.Equal(t, res, "test")
}

func Test_goRedis_Set_NX(t *testing.T) {
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
}

func Test_goRedis_Set_XX(t *testing.T) {
	key := "Test_goRedis_SetXX"
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
}

func Test_goRedis_Set_EX(t *testing.T) {
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
}

func Test_goRedis_Set_EX_NX(t *testing.T) {
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
}

func Test_goRedis_Set_EX_XX(t *testing.T) {
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
}

func Test_goRedis_Set_PX(t *testing.T) {
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
}

func Test_goRedis_Set_PX_NX(t *testing.T) {
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
}

func Test_goRedis_Set_PX_XX(t *testing.T) {
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
}

func Test_goRedis_Subscribe(t *testing.T) {
	channel := "Test_goRedis_Subscribe"
	msg := make(chan *redis.Message)
	New("127.0.0.1", 6379).Subscribe([]string{channel}, func(sub interface{}) {
		for {
			select {
			case msg <- <-sub.(*redis.PubSub).Channel():
				sub.(*redis.PubSub).Close()
			}
		}
	})

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
		(<-msg).Payload,
		message,
	)
}
