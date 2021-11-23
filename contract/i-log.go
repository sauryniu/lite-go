package contract

type ILog interface {
	AddLabel(key, format string, v ...interface{}) ILog
	Debug()
	Error(err error)
	Fatal()
	Info()
	Warning()
}
