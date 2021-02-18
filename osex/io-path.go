package osex

import (
	"path/filepath"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/ioex"
)

type osPath struct{}

func (m osPath) Join(paths ...string) string {
	var res string
	underscore.Chain(paths).Aggregate(func(memo string, r string, _ int) string {
		if memo == "" {
			return r
		}

		if r == ".." {
			return filepath.Dir(memo)
		}

		return filepath.Join(memo, r)
	}, "").Value(&res)
	return res
}

// NewIOPath is 路径实例
func NewIOPath() ioex.IPath {
	return new(osPath)
}
