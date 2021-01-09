package mongodb

import (
	"testing"

	"github.com/ahl5esoft/lite-go/dp/ioc"
	"github.com/stretchr/testify/assert"
)

func Test_NewStartup(t *testing.T) {
	err := NewStartup().Handle(&StartupContext{
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
