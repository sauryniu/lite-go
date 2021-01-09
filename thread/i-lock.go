package thread

import "time"

// ILock is 锁接口
type ILock interface {
	Lock(string, ...interface{}) (bool, error)
	SetExpire(seconds time.Duration) ILock
	Unlock()
}
