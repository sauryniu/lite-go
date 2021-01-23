package mongodb

import (
	"github.com/ahl5esoft/lite-go/object"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type stringGenerator struct{}

func (m stringGenerator) Generate() string {
	return primitive.NewObjectID().Hex()
}

// NewStringGenerator is 字符串生成器
func NewStringGenerator() object.IStringGenerator {
	return new(stringGenerator)
}
