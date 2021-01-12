package thread

import "time"

// LockIoCKey is 依赖注入键
const LockIoCKey = "lock"

// ILock is 锁接口
type ILock interface {
	Lock(string, ...interface{}) (bool, error)
	SetExpire(seconds time.Duration) ILock
	Unlock()
}
