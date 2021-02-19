package log

// ILog is 日志接口
type ILog interface {
	AddAttr(key, format string, v ...interface{}) ILog
	Debug()
	Error()
	Info()
	Warning()
}
