package ginex

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
)

// NewPostRunOption is post请求运行选项
func NewPostRunOption() RunOption {
	return func(app *gin.Engine) {
		app.POST(api.RouteRule, func(ctx *gin.Context) {
			var rp api.RouteParam
			ctx.ShouldBindUri(&rp)

			var resp api.Response
			resp.Data = ""

			var err error
			defer func() {
				if rv := recover(); rv != nil {
					if cErr, ok := rv.(error); ok {
						err = cErr
					} else {
						err = fmt.Errorf("%v", rv)
					}
				}

				if err != nil {
					if cErr, ok := err.(api.CustomError); ok {
						resp.Error = cErr.Code
						resp.Data = cErr.Error()
					} else {
						fmt.Println(err)
						debug.PrintStack()
						resp.Error = api.PanicErrorCode
					}
				}

				ctx.JSON(http.StatusOK, resp)
			}()

			apiInstance := api.New(rp.Endpoint, rp.Name)
			if err = ctx.ShouldBind(apiInstance); err != nil {
				err = api.NewError(api.VerifyErrorCode, "")
				return
			}

			resp.Data, err = apiInstance.Call()
		})
	}
}
