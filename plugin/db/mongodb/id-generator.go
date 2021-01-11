package mongodb

import (
	"github.com/ahl5esoft/lite-go/object"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type idGenerator struct{}

func (m idGenerator) Generate() string {
	return primitive.NewObjectID().Hex()
}

// NewIDGenerator is object.IIDGenerator实例
func NewIDGenerator() object.IIDGenerator {
	return new(idGenerator)
}
