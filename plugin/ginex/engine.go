package ginex

import (
	"fmt"
	"net/http"

	"github.com/ahl5esoft/lite-go/api"
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
			defer func() {
				if rv := recover(); rv != nil {
					if err, ok := rv.(api.CustomError); ok {
						resp.Error = int(err.Code)
						resp.Data = err.Error()
					} else {
						fmt.Println(rv)

						resp.Data = ""
						resp.Error = int(api.PanicErrorCode)
					}
				}

				ctx.JSON(http.StatusOK, resp)
			}()

			var us uriStruct
			if err := ctx.ShouldBindUri(&us); err != nil {
				panic(err)
			}

			a := api.New(us.Endpoint, us.Name)
			if !a.Auth() {
				resp.Data = ""
				resp.Error = int(api.AuthErrorCode)
			} else if !a.Valid(ctx) {
				resp.Data = ""
				resp.Error = int(api.VerifyErrorCode)
			} else {
				resp.Data = a.Call()
			}
		},
	)
	return app
}
