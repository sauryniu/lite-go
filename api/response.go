package api

import "github.com/ahl5esoft/lite-go/errorex"

// Response is api响应结构
type Response struct {
	Data  interface{}  `json:"data"`
	Error errorex.Code `json:"err"`
}
