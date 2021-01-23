package goredis

import (
	"github.com/ahl5esoft/lite-go/plugin/redisex"
	"github.com/go-redis/redis"
)

// NewDatabaseOption is 数据库选项
func NewDatabaseOption(database int) redisex.Option {
	return func(opt interface{}) {
		opt.(*redis.Options).DB = database
	}
}

// NewPasswordOption is 密码选项
func NewPasswordOption(password string) redisex.Option {
	return func(opt interface{}) {
		opt.(*redis.Options).Password = password
	}
}
