package ginex

import (
	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/gin-gonic/gin"
)

// RunOption is 运行选项
type RunOption = func(*gin.Engine)

// Run is 运行
func Run(mode string, options ...RunOption) {
	gin.SetMode(mode)

	app := gin.New()
	underscore.Chain(options).Each(func(r RunOption, _ int) {
		r(app)
	})
}
