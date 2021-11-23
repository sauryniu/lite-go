package ginsvc

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewPortOpion(port int) Option {
	return func(app *gin.Engine) {
		addr := fmt.Sprintf(":%d", port)
		fmt.Println(addr)
		err := app.Run(addr)
		if err != nil {
			panic(err)
		}
	}
}
