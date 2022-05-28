package mongosvc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type driverFactory struct {
	Ctx context.Context

	name   string
	client *mongo.Client
	db     *mongo.Database
	option *options.ClientOptions
}

func (m *driverFactory) BuildClient() (*mongo.Client, error) {
	if m.client == nil {
		var err error
		if m.client, err = mongo.Connect(m.Ctx, m.option); err != nil {
			return nil, err
		}
	}

	return m.client, nil
}

func (m *driverFactory) BuildCollection(model *modelMetadata) (*mongo.Collection, error) {
	name, err := model.GetTableName()
	if err != nil {
		return nil, err
	}

	db, err := m.BuildDb()
	if err != nil {
		return nil, err
	}

	return db.Collection(name), nil
}

func (m *driverFactory) BuildDb() (*mongo.Database, error) {
	if m.db == nil {
		client, err := m.BuildClient()
		if err != nil {
			return nil, err
		}

		m.db = client.Database(m.name)
	}

	return m.db, nil
}

func newDriverFactory(name, uri string) *driverFactory {
	return &driverFactory{
		Ctx:    context.Background(),
		name:   name,
		option: options.Client().ApplyURI(uri),
	}
}
