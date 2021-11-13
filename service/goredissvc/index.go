package goredissvc

import (
	"context"
	"sync"
	"time"

	"github.com/ahl5esoft/lite-go/contract"
	"github.com/go-redis/redis/v8"
)

var adapterMutex sync.Mutex

type adapter struct {
	cfg    *redis.Options
	client redis.Cmdable
	ctx    context.Context
}

func (m *adapter) Get(k string) (string, error) {
	res, err := m.getClient().Get(m.ctx, k).Result()
	if err != nil && err == redis.Nil {
		return "", nil
	}

	return res, err
}

func (m *adapter) Set(k, v string, expires time.Duration) (bool, error) {
	res, err := m.getClient().Set(m.ctx, k, v, expires).Result()
	if err != nil {
		return false, err
	}

	return res == "OK", nil
}

func (m *adapter) WithContext(ctx context.Context) interface{} {
	return &adapter{
		client: m.getClient(),
		ctx:    ctx,
		cfg:    m.cfg,
	}
}

func (m *adapter) getClient() redis.Cmdable {
	if m.client == nil {
		adapterMutex.Lock()
		defer adapterMutex.Unlock()

		if m.client == nil {
			m.client = redis.NewClient(m.cfg)
		}
	}
	return m.client
}

func NewRedis(cfg redis.Options) contract.IRedis {
	return &adapter{
		ctx: context.Background(),
		cfg: &cfg,
	}
}
