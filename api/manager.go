package api

var metadatas = make(map[string]map[string]func() IAPI)

// New is 创建api实例
func New(endpoint, name string) IAPI {
	if ctors, ok := metadatas[endpoint]; ok {
		if ctor, ok := ctors[name]; ok {
			return ctor()
		}
	}

	return invalid
}

// Register is 注册api
func Register(endpoint, name string, ctor func() IAPI) {
	if _, ok := metadatas[endpoint]; !ok {
		metadatas[endpoint] = make(map[string]func() IAPI)
	}

	metadatas[endpoint][name] = ctor
}
