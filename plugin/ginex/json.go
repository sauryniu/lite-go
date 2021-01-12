package ginex

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Data  interface{}   `json:"data"`
	Error api.ErrorCode `json:"err"`
}

// JSON is 输出json
func JSON(err error, apiInstance api.IAPI, ctx *gin.Context) {
	var resp jsonResponse
	resp.Data = ""
	defer func() {
		if rv := recover(); rv != nil {
			if cErr, ok := rv.(error); ok {
				err = cErr
			} else {
				fmt.Println(err)
				debug.PrintStack()
				resp.Error = api.PanicErrorCode
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

	if err == nil {
		if apiInstance != nil {
			resp.Data, err = apiInstance.Call()
		} else {
			err = api.NewError(api.APIErrorCode, "")
		}
	}
}
