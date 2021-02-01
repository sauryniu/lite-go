//go:generate mockgen -destination i-api-caller_mock.go -package session github.com/ahl5esoft/lite-go/api/session IAPICaller

package session

import "time"

// IAPICaller is 会话调用者
type IAPICaller interface {
	Get(k string, v interface{}) error
	Set(body interface{}, expires, interval time.Duration) (string, error)
}
