package ginex

import (
	"github.com/ahl5esoft/lite-go/api"
	"github.com/gin-gonic/gin"
)

// API is gin实现api.IAPI
type API struct {
	Ctx     *gin.Context `binding:"-"`
	Derived api.IAPI     `binding:"-"`

	verifyError error

	AuthFunc         func() bool
	CallFunc         func() (interface{}, error)
	SetRequestAction func(req interface{})
}

// Auth is 认证
func (m API) Auth() bool {
	if m.AuthFunc != nil {
		return m.AuthFunc()
	}

	return true
}

// Call is 调用
func (m API) Call() (interface{}, error) {
	return m.CallFunc()
}

// SetRequest is 设置请求
func (m *API) SetRequest(req interface{}) {
	m.Ctx = req.(*gin.Context)
	if m.Derived != nil {
		m.verifyError = m.Ctx.ShouldBind(m.Derived)
	}

	if m.SetRequestAction != nil {
		m.SetRequestAction(req)
	}
}

// Valid is 验证
func (m API) Valid() bool {
	return m.verifyError == nil
}
