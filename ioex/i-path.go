//go:generate mockgen -destination i-path_mock.go -package ioex github.com/ahl5esoft/lite-go/ioex IPath

package ioex

// IPath is 路径接口
type IPath interface {
	Join(paths ...string) string
}
