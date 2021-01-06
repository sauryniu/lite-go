package ioos

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/ahl5esoft/lite-go/lib/ioex"
)

type file struct {
	ioex.INode
}

func (m file) GetExt() string {
	filePath := m.GetPath()
	return filepath.Ext(filePath)
}

func (m file) GetFile() (*os.File, error) {
	var file *os.File
	var err error
	filePath := m.GetPath()
	if m.IsExist() {
		file, err = os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	} else {
		file, err = os.Create(filePath)
	}
	return file, err
}

func (m file) Read(data interface{}) error {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Ptr {
		return fmt.Errorf("osex.file.Read: data必须为指针")
	}

	f, err := m.GetFile()
	if err != nil {
		return err
	}

	defer f.Close()

	bf, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	value = value.Elem()
	if value.Kind() == reflect.String {
		value.SetString(
			string(bf),
		)
		return nil
	} else if value.Kind() == reflect.Slice && value.Type().Elem().Kind() == reflect.Uint8 {
		value.SetBytes(bf)
		return nil
	}

	return fmt.Errorf(
		"osex.file.Read: 不支持%s",
		value.Type(),
	)
}

func (m file) Write(data interface{}) error {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.String {
		return m.writeString(
			data.(string),
		)
	}

	return fmt.Errorf("osex.file.Write暂不支持%s", dataType.Kind())
}

func (m file) writeString(s string) error {
	file, err := m.GetFile()
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(s)
	return err
}

// NewFile is 创建io.IFile实例
func NewFile(pathArgs ...string) ioex.IFile {
	return &file{
		INode: newNode(pathArgs...),
	}
}
