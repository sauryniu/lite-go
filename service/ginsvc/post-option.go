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
	"github.com/ahl5esoft/lite-go/service/errorsvc"
	"github.com/ahl5esoft/lite-go/service/iocsvc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	jsoniter "github.com/json-iterator/go"
)

func NewPostOption(apiFactory contract.IApiFactory) Option {
	return func(app *gin.Engine) {
		validate := validator.New()
		app.POST("/:endpoint/:api", func(ctx *gin.Context) {
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

			api := apiFactory.Build(
				ctx.Param("endpoint"),
				ctx.Param("api"),
			)

			if ctx.Request.ContentLength > 0 {
				var bodyBytes []byte
				if bodyBytes, err = ioutil.ReadAll(ctx.Request.Body); err != nil {
					return
				}

				ctx.Set(contextkey.HttpBody, bodyBytes)

				if err = jsoniter.Unmarshal(bodyBytes, &api); err != nil {
					return
				}

				if err = validate.Struct(api); err != nil {
					err = errorsvc.Newf(errorcode.Verify, "")
					return
				}
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
