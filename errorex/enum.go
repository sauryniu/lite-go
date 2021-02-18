package errorex

// Code is 错误码
type Code int

const (
	// APICode is api错误码
	APICode Code = iota + 501
	// AuthCode is 认证错误码
	AuthCode
	// VerifyCode is 验证错误码
	VerifyCode
	// TipCode is 提醒错误码
	TipCode
	// PanicCode is 异常错误码
	PanicCode Code = 599
)
