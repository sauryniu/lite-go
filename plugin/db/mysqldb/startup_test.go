package mysqldb

import (
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/stretchr/testify/assert"
)

type startupContext struct {
	MysqlOption FactoryOption
}

func (m startupContext) GetMysqlOption() FactoryOption {
	return m.MysqlOption
}

func Test_NewStartupHandler(t *testing.T) {
	err := NewStartupHandler().Handle(&startupContext{
		MysqlOption: FactoryOption{
			DbName:   "go-test",
			Host:     "10.1.33.67",
			Password: "123456",
			Username: "root",
		},
	})
	assert.NoError(t, err)
	assert.True(
		t,
		ioc.Has("db"),
	)
}
