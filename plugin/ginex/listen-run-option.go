package ginex

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// NewListenRunOption is 监听运行选项
func NewListenRunOption(port int) RunOption {
	return func(app *gin.Engine) {
		addr := fmt.Sprintf(":%d", port)
		fmt.Println(addr)
		app.Run(addr)
	}
}
