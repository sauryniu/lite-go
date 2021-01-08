package ginex

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
)

type publicAPI struct {
}

func (m publicAPI) Auth() bool {
	return true
}

func (m *publicAPI) ValidDerive(ctx interface{}, derive api.IAPI) bool {
	err := ctx.(*gin.Context).ShouldBind(derive)
	if err != nil {
		fmt.Println("ValidDerive", err)
	}
	return err == nil
}
