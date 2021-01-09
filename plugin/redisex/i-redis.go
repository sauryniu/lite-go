package redisex

// IRedis is redis接口
type IRedis interface {
	Del(...string) (int, error)
	Exists(string) (bool, error)
	Get(string) (string, error)
	Set(string, string, ...interface{}) (bool, error)
}
