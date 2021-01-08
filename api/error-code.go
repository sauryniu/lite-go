package api

// ErrorCode is 错误码
type ErrorCode int

const (
	// APIErrorCode is api错误码
	APIErrorCode ErrorCode = iota + 500
	// AuthErrorCode is 认证错误码
	AuthErrorCode
	// VerifyErrorCode is 验证错误码
	VerifyErrorCode
	// PanicErrorCode is 异常错误码
	PanicErrorCode ErrorCode = 599
)
