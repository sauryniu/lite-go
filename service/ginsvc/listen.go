package ginsvc

import "github.com/gin-gonic/gin"

func Listen(options ...Option) {
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	for _, r := range options {
		r(app)
	}
}
