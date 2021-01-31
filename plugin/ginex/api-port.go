package ginex

import (
	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
)

type apiPort []APIPortOption

func (m apiPort) Listen() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	underscore.Chain(m).Each(func(r APIPortOption, _ int) {
		r(app)
	})
}

// NewAPIPort is 创建api端口(gin实现)
func NewAPIPort(options ...APIPortOption) api.IPort {
	return apiPort(options)
}
