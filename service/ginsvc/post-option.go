package ginsvc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/ahl5esoft/lite-go/contract"
	contextkey "github.com/ahl5esoft/lite-go/model/enum/context-key"
	errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"
	"github.com/ahl5esoft/lite-go/model/response"
	"github.com/ahl5esoft/lite-go/service/iocsvc"

	"github.com/gin-gonic/gin"
)

func NewPostOption(
	relativePath string,
	getApiFunc func(ctx *gin.Context) (contract.IApi, error),
) Option {
	return func(app *gin.Engine) {
		app.POST(relativePath, func(ctx *gin.Context) {
			var resp response.Api
			defer func() {
				ctx.JSON(http.StatusOK, resp)
			}()

			var err error
			defer func() {
				if rv := recover(); rv != nil {
					var ok bool
					if err, ok = rv.(error); !ok {
						err = fmt.Errorf("%v", rv)
					}
				}

				if err != nil {
					if cErr, ok := err.(contract.IError); ok {
						resp.Error = cErr.GetCode()
						if cErr.GetData() != nil {
							resp.Data = cErr.GetData()
						} else {
							resp.Data = cErr.Error()
						}
					} else {
						resp.Data = err.Error()
						resp.Error = errorcode.Panic
					}
				}
			}()

			if ctx.Request.ContentLength > 0 {
				var bodyBytes []byte
				if bodyBytes, err = ioutil.ReadAll(ctx.Request.Body); err != nil {
					return
				}

				ctx.Set(contextkey.HttpBody, bodyBytes)
			}

			var api contract.IApi
			if api, err = getApiFunc(ctx); err != nil {
				return
			}

			iocsvc.Inject(api, func(v reflect.Value) reflect.Value {
				if w, ok := v.Interface().(contract.IContextWrapper); ok {
					return reflect.ValueOf(
						w.WithContext(ctx),
					)
				}
				return v
			})

			resp.Data, err = api.Call()
		})
	}
}
