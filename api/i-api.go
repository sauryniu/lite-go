//go:generate mockgen -destination i-api_mock.go -package api github.com/ahl5esoft/lite-go/api IAPI

package api

// IAPI is api接口
type IAPI interface {
	Call() (interface{}, error)
}
