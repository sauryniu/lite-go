package contract

import errorcode "github.com/ahl5esoft/lite-go/model/enum/error-code"

type IError interface {
	error

	GetCode() errorcode.Value
	GetData() interface{}
}
