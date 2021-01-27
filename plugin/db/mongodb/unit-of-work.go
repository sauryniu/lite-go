package mongodb

import (
	"reflect"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/plugin/db/identity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type unitOfWork struct {
	addQueue    []identity.IIdentity
	pool        *connectPool
	removeQueue []identity.IIdentity
	saveQueue   []identity.IIdentity
}

func (m *unitOfWork) Commit() error {
	client, err := m.pool.GetClient()
	if err != nil {
		return err
	}

	return client.UseSession(m.pool.ctx, func(ctx mongo.SessionContext) (err error) {
		if err = m.commitAdd(ctx); err != nil {
			return err
		}

		if err = m.commitRemove(ctx); err != nil {
			return err
		}

		return m.commitSave(ctx)
	})
}

func (m *unitOfWork) commitAdd(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.addQueue).Map(func(r identity.IIdentity, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		s := identity.NewStruct(
			entryValue.Type(),
		)
		col, rErr := m.pool.GetCollection(s)
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
	m.addQueue = make([]identity.IIdentity, 0)
	return
}

func (m *unitOfWork) commitRemove(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.removeQueue).Map(func(r identity.IIdentity, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		s := identity.NewStruct(
			entryValue.Type(),
		)
		col, rErr := m.pool.GetCollection(s)
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
	m.removeQueue = make([]identity.IIdentity, 0)
	return
}

func (m *unitOfWork) commitSave(ctx mongo.SessionContext) (err error) {
	underscore.Chain(m.saveQueue).Map(func(r identity.IIdentity, _ int) error {
		entryValue := reflect.ValueOf(r)
		if entryValue.Kind() == reflect.Ptr {
			entryValue = entryValue.Elem()
		}
		s := identity.NewStruct(
			entryValue.Type(),
		)
		col, rErr := m.pool.GetCollection(s)
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
	m.saveQueue = make([]identity.IIdentity, 0)
	return
}

func (m *unitOfWork) registerAdd(entry identity.IIdentity) {
	m.addQueue = append(m.addQueue, entry)
}

func (m *unitOfWork) registerRemove(entry identity.IIdentity) {
	m.removeQueue = append(m.removeQueue, entry)
}

func (m *unitOfWork) registerSave(entry identity.IIdentity) {
	m.saveQueue = append(m.saveQueue, entry)
}

func newUnitOfWork(pool *connectPool) *unitOfWork {
	return &unitOfWork{
		addQueue:    make([]identity.IIdentity, 0),
		pool:        pool,
		removeQueue: make([]identity.IIdentity, 0),
		saveQueue:   make([]identity.IIdentity, 0),
	}
}
