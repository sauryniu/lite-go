package object

// StringGeneratorIoCKey is IStringGenerator依赖注入键
const StringGeneratorIoCKey = "string-generator"

// IStringGenerator is 字符串生成接口
type IStringGenerator interface {
	Generate() string
}
