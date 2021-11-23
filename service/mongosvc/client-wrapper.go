package mongosvc

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientWrapperMutex sync.Mutex

type clientWrapper struct {
	Ctx context.Context

	client *mongo.Client
	db     *mongo.Database
	option *options.ClientOptions
	name   string
}

func (m *clientWrapper) GetClient() (*mongo.Client, error) {
	if m.client == nil {
		clientWrapperMutex.Lock()
		defer clientWrapperMutex.Unlock()

		if m.client == nil {
			var err error
			if m.client, err = mongo.Connect(m.Ctx, m.option); err != nil {
				return nil, err
			}
		}
	}

	return m.client, nil
}

func (m *clientWrapper) GetDb() (*mongo.Database, error) {
	if m.db == nil {
		client, err := m.GetClient()
		if err != nil {
			return nil, err
		}

		clientWrapperMutex.Lock()
		defer clientWrapperMutex.Unlock()

		if m.db == nil {
			m.db = client.Database(m.name)
		}
	}

	return m.db, nil
}

func (m *clientWrapper) GetCollection(model *modelMetadata) (*mongo.Collection, error) {
	name, err := model.GetTableName()
	if err != nil {
		return nil, err
	}

	db, err := m.GetDb()
	if err != nil {
		return nil, err
	}

	return db.Collection(name), nil
}

func newClientWrapper(name, uri string) *clientWrapper {
	return &clientWrapper{
		Ctx:    context.Background(),
		name:   name,
		option: options.Client().ApplyURI(uri),
	}
}
