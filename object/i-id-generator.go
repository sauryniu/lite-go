package object

// IIDGenerator is id生成接口
type IIDGenerator interface {
	Generate() string
}
