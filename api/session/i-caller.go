//go:generate mockgen -destination i-caller_mock.go -package session github.com/ahl5esoft/lite-go/api/session ICaller

package session

import "time"

// ICaller is 会话调用者
type ICaller interface {
	Get(k string, v interface{}) error
	Set(body interface{}, expires, interval time.Duration) (string, error)
}
