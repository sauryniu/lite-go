package osex

import "time"

// ICommand is 命令接口
type ICommand interface {
	Exec(name string, args ...string) (stdout string, stderr string, err error)
	SetDir(format string, args ...interface{}) ICommand
	SetExpires(expires time.Duration) ICommand
}
