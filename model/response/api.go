package response

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

type Api struct {
	Data  interface{}     `json:"data"`
	Error errorcode.Value `json:"err"`
}
