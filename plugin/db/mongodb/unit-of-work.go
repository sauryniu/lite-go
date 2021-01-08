package mongodb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type unitOfWork struct {
	AddQueue    []identity.IIdentity
	Pool        *connectPool
	RemoveQueue []identity.IIdentity
	SaveQueue   []identity.IIdentity
}

func (m *unitOfWork) Commit() error {
	client, err := m.Pool.GetClient()
	if err != nil {
		return err
	}

	return client.UseSession(m.Pool.Ctx, func(ctx mongo.SessionContext) (err error) {
		if err = m.commitAdd(ctx); err != nil {
			return err
		}

		if err = m.commitRemove(ctx); err != nil {
			return err
		}

		return m.commitSave(ctx)
	})
}

func (m *unitOfWork) RegisterAdd(entry identity.IIdentity) {
	m.AddQueue = append(m.AddQueue, entry)
}

func (m *unitOfWork) RegisterRemove(entry identity.IIdentity) {
	m.RemoveQueue = append(m.RemoveQueue, entry)
}

func (m *unitOfWork) RegisterSave(entry identity.IIdentity) {
	m.SaveQueue = append(m.SaveQueue, entry)
}

func (m *unitOfWork) commitAdd(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.AddQueue).Map(func(r identity.IIdentity, _ int) error {
		entryValue := reflect.ValueOf(r)
		s := identity.NewStruct(
			entryValue.Type(),
		)
		col, rErr := m.Pool.GetCollection(s)
		if rErr != nil {
			return rErr
		}

		doc := make(bson.M)
		underscore.Chain(
			s.FindFields(),
		).Each(func(r identity.IField, _ int) {
			doc[r.GetName()] = r.GetValue(entryValue)
		})
		_, rErr = col.InsertOne(ctx, doc)
		return rErr
	}).Find(func(r error, _ int) bool {
		return r != nil
	}).Value(&err)
	m.AddQueue = make([]identity.IIdentity, 0)
	return
}

func (m *unitOfWork) commitRemove(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.RemoveQueue).Map(func(r identity.IIdentity, _ int) error {
		entryValue := reflect.ValueOf(r)
		s := identity.NewStruct(
			entryValue.Type(),
		)
		col, rErr := m.Pool.GetCollection(s)
		if rErr != nil {
			return rErr
		}

		idField, rErr := s.GetIDField()
		if rErr != nil {
			return rErr
		}

		filter := make(bson.M)
		filter[idField.GetName()] = idField.GetValue(entryValue)
		_, rErr = col.DeleteOne(ctx, filter)
		return rErr
	}).Find(func(r error, _ int) bool {
		return r != nil
	}).Value(&err)
	m.RemoveQueue = make([]identity.IIdentity, 0)
	return
}

func (m *unitOfWork) commitSave(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.SaveQueue).Map(func(r identity.IIdentity, _ int) error {
		entryValue := reflect.ValueOf(r)
		s := identity.NewStruct(
			entryValue.Type(),
		)
		col, rErr := m.Pool.GetCollection(s)
		if rErr != nil {
			return rErr
		}

		filer := make(bson.M)
		doc := make(bson.M)
		underscore.Chain(
			s.FindFields(),
		).Each(func(r identity.IField, _ int) {
			if r.GetStructName() != "" {
				filer[r.GetName()] = r.GetValue(entryValue)
			}
			doc[r.GetName()] = r.GetValue(entryValue)
		})
		_, rErr = col.UpdateOne(ctx, filer, bson.M{
			"$set": doc,
		})
		return rErr
	}).Find(func(r error, _ int) bool {
		return r != nil
	}).Value(&err)
	m.SaveQueue = make([]identity.IIdentity, 0)
	return
}

func newUnitOfWork(pool *connectPool) *unitOfWork {
	return &unitOfWork{
		AddQueue:    make([]identity.IIdentity, 0),
		Pool:        pool,
		RemoveQueue: make([]identity.IIdentity, 0),
		SaveQueue:   make([]identity.IIdentity, 0),
	}
}
