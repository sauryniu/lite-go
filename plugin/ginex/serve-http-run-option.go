package ginex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewServeHTTPRunOption is http运行选项
func NewServeHTTPRunOption(req *http.Request, resp http.ResponseWriter) RunOption {
	return func(app *gin.Engine) {
		app.ServeHTTP(resp, req)
	}
}
