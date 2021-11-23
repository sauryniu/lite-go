package mongosvc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_clientWrapper_GetClient(t *testing.T) {
	client, err := testClient.GetClient()
	assert.NoError(t, err)

	err = client.Ping(
		context.Background(),
		nil,
	)
	assert.NoError(t, err)
}
