package api

import "github.com/ahl5esoft/lite-go/errorex"

type invalidAPI struct{}

func (m invalidAPI) Call() (interface{}, error) {
	return nil, errorex.New(errorex.APICode, "")
}
