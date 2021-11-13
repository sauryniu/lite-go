package contract

import "time"

type IRedis interface {
	Get(k string) (string, error)
	Set(k, v string, expires time.Duration) (bool, error)
}
