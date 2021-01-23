package thread

import "time"

// LockIoCKey is 依赖注入键
const LockIoCKey = "lock"

// ILock is 锁接口
type ILock interface {
	Lock(format string, args ...interface{}) (unlock func(), err error)
	SetExpire(seconds time.Duration) ILock
}
