package osex

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/log"
)

type osLog struct {
	attr map[string]string
}

func (m *osLog) AddAttr(key, format string, v ...interface{}) log.ILog {
	m.attr[key] = fmt.Sprintf(format, v...)
	return m
}

func (m *osLog) AddDesc(format string, v ...interface{}) log.ILog {
	m.AddAttr("desc", format, v...)
	return m
}

func (m *osLog) Debug() {
	m.print("debug")
}

func (m *osLog) Error() {
	m.print("error")
}

func (m *osLog) Info() {
	m.print("info")
}

func (m *osLog) Warning() {
	m.print("warning")
}

func (m *osLog) print(t string) {
	m.AddAttr("type", t)
	fmt.Println(m.attr)
	m.attr = make(map[string]string)
}

// NewLog is 创建log.ILog实例
func NewLog() log.ILog {
	return &osLog{
		attr: make(map[string]string),
	}
}
