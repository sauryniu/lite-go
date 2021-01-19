package ioos

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_file_GetExt(t *testing.T) {
	res := NewFile("a.txt").GetExt()
	assert.Equal(t, res, ".txt")
}

func Test_file_GetFile(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	file := NewFile(wd, "get-file.txt")
	f, err := file.GetFile()
	assert.NoError(t, err)

	err = f.Close()
	assert.NoError(t, err)

	err = file.Remove()
	assert.NoError(t, err)
}

func Test_file_Read_Bytes(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	file := NewFile(wd, "read.txt")
	defer file.Remove()

	f, err := file.GetFile()
	assert.NoError(t, err)

	defer f.Close()

	text := "read string"
	_, err = f.WriteString(text)
	assert.NoError(t, err)

	var res []byte
	err = file.Read(&res)
	assert.NoError(t, err)
	assert.Equal(
		t,
		string(res),
		text,
	)
}

func Test_file_Read_JSON(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	file := NewFile(wd, "read.txt")
	defer file.Remove()

	f, err := file.GetFile()
	assert.NoError(t, err)

	defer f.Close()

	text := `{"name":"n","age":11}`
	_, err = f.WriteString(text)
	assert.NoError(t, err)

	type testStruct struct {
		Name string
		Age  int
	}
	var v testStruct
	err = file.Read(&v)
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		"不支持ioos.file.Read(ioos.testStruct)",
	)
}

func Test_file_Read_String(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	file := NewFile(wd, "read.txt")
	defer file.Remove()

	f, err := file.GetFile()
	assert.NoError(t, err)

	defer f.Close()

	text := "read string"
	_, err = f.WriteString(text)
	assert.NoError(t, err)

	var res string
	err = file.Read(&res)
	assert.NoError(t, err)
	assert.Equal(t, res, text)
}

func Test_file_ReadJSON(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	file := NewFile(wd, "read-json.txt")
	defer file.Remove()

	f, err := file.GetFile()
	assert.NoError(t, err)

	defer f.Close()

	text := `{"name":"n","age":11}`
	_, err = f.WriteString(text)
	assert.NoError(t, err)

	type testStruct struct {
		Name string
		Age  int
	}
	var v testStruct
	err = file.ReadJSON(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, testStruct{
		Name: "n",
		Age:  11,
	})
}

func Test_file_ReadYaml(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	file := NewFile(wd, "read-yaml.txt")
	defer file.Remove()

	f, err := file.GetFile()
	assert.NoError(t, err)

	defer f.Close()

	text := `name: n
age: 11`
	_, err = f.WriteString(text)
	assert.NoError(t, err)

	type testStruct struct {
		Name string
		Age  int
	}
	var v testStruct
	err = file.ReadYaml(&v)
	assert.NoError(t, err)
	assert.Equal(t, v, testStruct{
		Name: "n",
		Age:  11,
	})
}

func Test_file_Write_String(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	file := NewFile(wd, "write.txt")
	err = file.Write("aa")

	defer file.Remove()

	assert.NoError(t, err)

	f, err := file.GetFile()
	assert.NoError(t, err)

	defer f.Close()

	res, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(
		t,
		string(res),
		"aa",
	)
}
