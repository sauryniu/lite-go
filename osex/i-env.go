package osex

// IoCKey is 依赖注入键
const IoCKey = "env"

// IEnv is 环境变量接口
type IEnv interface {
	Get(string, interface{})
}
