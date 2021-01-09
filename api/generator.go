package api

import (
	"github.com/ahl5esoft/lite-go/ioex"
)

// GenerateOption is 生成选项
type GenerateOption struct {
	SrcDir ioex.IDirectory
	DstDir ioex.IDirectory
}

// Generate is 生成API元文件
func Generate(opt GenerateOption) {

}
