package session

import (
	"time"

	"github.com/ahl5esoft/lite-go/api"
	jsoniter "github.com/json-iterator/go"
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

	getRoute, setRoute string
}

func (m apiCaller) Get(k string, v interface{}) error {
	res, err := m.ICaller.Call(m.getRoute, getMessage{
		Key: k,
	})
	if err != nil {
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

	res, err := m.ICaller.Call(m.setRoute, setMessage{
		Expires:  int(expires / time.Second),
		Interval: intervalCount,
		Value:    bodyJSON,
	})
	if err != nil {
		return "", err
	}

	return res.(string), nil
}

// NewAPICaller is 创建会话调用
func NewAPICaller(caller api.ICaller, getRoute, setRoute string) IAPICaller {
	return &apiCaller{
		ICaller:  caller,
		getRoute: getRoute,
		setRoute: setRoute,
	}
}
