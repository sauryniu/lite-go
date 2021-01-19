package osex

import (
	"os"
	"strings"

	"github.com/ahl5esoft/lite-go/ioex"
)

type projectEnv struct {
	IEnv

	project string
}

func (m projectEnv) Get(k string, v interface{}) {
	k = strings.Join([]string{m.project, k}, "-")
	k = strings.Replace(k, "-", "_", -1)
	m.IEnv.Get(k, v)
}

// NewProjectEnv is 项目env
func NewProjectEnv(ioFactory ioex.IFactory, env IEnv) IEnv {
	wd, _ := os.Getwd()
	return &projectEnv{
		IEnv:    env,
		project: ioFactory.BuildDirectory(wd).GetName(),
	}
}
