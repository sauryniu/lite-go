package osex

// IEnv is 环境变量接口
type IEnv interface {
	Get(string, interface{})
}
