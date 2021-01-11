package mongodb

import (
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/stretchr/testify/assert"
)

type startupContext struct {
	MongoOption FactoryOption
}

func (m startupContext) GetMongoOption() FactoryOption {
	return m.MongoOption
}

func Test_NewStartupHandler(t *testing.T) {
	err := NewStartupHandler().Handle(&startupContext{
		MongoOption: FactoryOption{
			DbName: "lite-go",
			URI:    "mongodb://localhost:27017",
		},
	})
	assert.NoError(t, err)

	assert.True(
		t,
		ioc.Has("db"),
	)
}
