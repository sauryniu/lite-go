package ginex

import (
	"fmt"
	"net/http"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/ahl5esoft/lite-go/errorex"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type routeParam struct {
	API      string `uri:"api"`
	Endpoint string `uri:"endpoint"`
}

// APIPortOption is api端口选项
type APIPortOption func(*gin.Engine)

// NewPostOption is post请求选项
func NewPostOption(apiFactory api.IFactory, validate *validator.Validate) APIPortOption {
	return func(app *gin.Engine) {
		app.POST("/:endpoint/:api", func(ctx *gin.Context) {
			var rp routeParam
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
					if cErr, ok := err.(errorex.Custom); ok {
						resp.Error = cErr.Code
						resp.Data = cErr.Error()
					} else {
						fmt.Println(
							fmt.Sprintf("%v", err),
						)
						resp.Error = errorex.PanicCode
					}
				}

				ctx.JSON(http.StatusOK, resp)
			}()

			apiInstance := apiFactory.Build(rp.Endpoint, rp.API)
			ctx.BindJSON(apiInstance)
			if err = validate.Struct(apiInstance); err != nil {
				err = errorex.New(errorex.VerifyCode, "")
				return
			}

			resp.Data, err = apiInstance.Call()
		})
	}
}

// NewRunOption is 监听运行选项
func NewRunOption(port int) APIPortOption {
	return func(app *gin.Engine) {
		addr := fmt.Sprintf(":%d", port)
		fmt.Println(addr)
		app.Run(addr)
	}
}

// NewServerHTTPOption is 服务http选项
func NewServerHTTPOption(req *http.Request, resp http.ResponseWriter) APIPortOption {
	return func(app *gin.Engine) {
		app.ServeHTTP(resp, req)
	}
}
