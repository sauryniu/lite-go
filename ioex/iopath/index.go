package iopath

import (
	"path/filepath"

	underscore "github.com/ahl5esoft/golang-underscore"
)

// Join is 拼接路径
func Join(pathArgs ...string) string {
	var res string
	underscore.Chain(pathArgs).Aggregate(func(memo string, r string, _ int) string {
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
