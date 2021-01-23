package randex

import (
	"math/rand"
	"time"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/object"
)

var defaultStringGeneratorSource = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

// NewStringGeneratorOption is 创建object.IStringGenerator选项
type NewStringGeneratorOption func(object.IStringGenerator)

type stringGenerator struct {
	length     int
	rand       *rand.Rand
	randSource rand.Source
	source     []byte
}

func (m stringGenerator) Generate() string {
	bf := make([]byte, m.length)
	for i := 0; i < len(bf); i++ {
		source := m.getSource()
		b := source[m.getRand().Intn(
			len(source),
		)]
		bf[i] = byte(b)
	}
	return string(bf)
}

func (m *stringGenerator) getRand() *rand.Rand {
	if m.rand == nil {
		randSource := m.getRandSource()
		m.rand = rand.New(randSource)
	}

	return m.rand
}

func (m *stringGenerator) getRandSource() rand.Source {
	if m.randSource == nil {
		seed := time.Now().UnixNano()
		m.randSource = rand.NewSource(seed)
	}

	return m.randSource
}

func (m *stringGenerator) getSource() []byte {
	if m.source == nil {
		m.source = defaultStringGeneratorSource
	}
	return m.source
}

// NewStringGenerator is 随机字符串生成器
func NewStringGenerator(length int, options ...NewStringGeneratorOption) object.IStringGenerator {
	generator := &stringGenerator{
		length: length,
	}
	underscore.Chain(options).Each(func(r NewStringGeneratorOption, _ int) {
		r(generator)
	})
	return generator
}

// NewStringGeneratorRandSourceOption is 随机源选项
func NewStringGeneratorRandSourceOption(randSource rand.Source) NewStringGeneratorOption {
	return func(generator object.IStringGenerator) {
		generator.(*stringGenerator).randSource = randSource
	}
}

// NewStringGeneratorSourceOption is 字符源选项
func NewStringGeneratorSourceOption(source []byte) NewStringGeneratorOption {
	return func(generator object.IStringGenerator) {
		generator.(*stringGenerator).source = source
	}
}
