package mongodb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connectPool_GetClient(t *testing.T) {
	client, err := pool.GetClient()
	assert.NoError(t, err)

	err = client.Ping(
		context.Background(),
		nil,
	)
	assert.NoError(t, err)
}
