package ginex

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/gin-gonic/gin"
)

type response struct {
	Data  interface{} `json:"data"`
	Error int         `json:"err"`
}

type uriStruct struct {
	Endpoint string `uri:"endpoint" binding:"required"`
	Name     string `uri:"name" binding:"required"`
}

func newEngine() *gin.Engine {
	app := gin.New()
	app.POST(
		"/:endpoint/:name",
		func(ctx *gin.Context) {
			var resp response
			resp.Data = ""

			var err error
			defer func() {
				if rv := recover(); rv != nil {
					if cErr, ok := rv.(api.CustomError); ok {
						resp.Error = int(cErr.Code)
						resp.Data = cErr.Error()
					} else {
						fmt.Println(rv)
						debug.PrintStack()
						resp.Error = int(api.PanicErrorCode)
					}
				} else if err != nil {
					fmt.Println(err)
					debug.PrintStack()
					resp.Error = int(api.PanicErrorCode)
				}

				ctx.JSON(http.StatusOK, resp)
			}()

			var us uriStruct
			if err = ctx.ShouldBindUri(&us); err != nil {
				return
			}

			a := api.New(us.Endpoint, us.Name)
			ioc.Inject(a)
			a.SetRequest(ctx)
			if !a.Auth() {
				resp.Error = int(api.AuthErrorCode)
			} else if !a.Valid() {
				resp.Error = int(api.VerifyErrorCode)
			} else {
				resp.Data, err = a.Call()
			}
		},
	)
	return app
}
