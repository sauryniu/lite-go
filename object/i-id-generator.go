package object

// IIDGeneratorIoCKey is IIDGenerator依赖注入键
const IIDGeneratorIoCKey = "id-generator"

// IIDGenerator is id生成接口
type IIDGenerator interface {
	Generate() string
}
