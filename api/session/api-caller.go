package session

import (
	"fmt"
	"time"

	"github.com/ahl5esoft/lite-go/api"
	jsoniter "github.com/json-iterator/go"
)

const (
	getRouteFormat = "%s/server/get"
	setRouteFormat = "%s/server/set"
)

type getMessage struct {
	Key string
}

type setMessage struct {
	Expires, Interval int
	Value             string
}

type apiCaller struct {
	api.ICaller

	app string
}

func (m apiCaller) Get(k string, v interface{}) error {
	res, err := m.ICaller.Call(
		fmt.Sprintf(getRouteFormat, m.app),
		getMessage{
			Key: k,
		},
		5*time.Second,
	)
	if err != nil || res.(string) == "" {
		return err
	}

	return jsoniter.UnmarshalFromString(
		res.(string),
		v,
	)
}

func (m apiCaller) Set(body interface{}, expires, interval time.Duration) (string, error) {
	bodyJSON, err := jsoniter.MarshalToString(body)
	if err != nil {
		return "", err
	}

	intervalCount := 0
	if interval > time.Second {
		intervalCount = int(interval / time.Second)
	}

	res, err := m.ICaller.Call(
		fmt.Sprintf(setRouteFormat, m.app),
		setMessage{
			Expires:  int(expires / time.Second),
			Interval: intervalCount,
			Value:    bodyJSON,
		},
		5*time.Second,
	)
	if err != nil {
		return "", err
	}

	return res.(string), nil
}

// NewAPICaller is 创建会话调用
func NewAPICaller(caller api.ICaller, app string) IAPICaller {
	return &apiCaller{
		ICaller: caller,
		app:     app,
	}
}
