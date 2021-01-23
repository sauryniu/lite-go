package thread

// LockIoCKey is 依赖注入键
const LockIoCKey = "lock"

// ILock is 锁接口
type ILock interface {
	Lock(key string, options ...LockOption) (unlock func(), err error)
}

// LockOption is 加锁选项
type LockOption func(ILock)
