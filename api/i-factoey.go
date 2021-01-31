//go:generate mockgen -destination i-factory_mock.go -package api github.com/ahl5esoft/lite-go/api IFactory

package api

// IFactory is api工厂接口
type IFactory interface {
	Build(endpoint, name string) IAPI
}
