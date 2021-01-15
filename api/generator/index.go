package generator

import (
	"html/template"
	"os"
	"regexp"
	"strings"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ahl5esoft/lite-go/ioex"
	"github.com/ahl5esoft/lite-go/ioex/ioos"
	"github.com/jessevdk/go-flags"
)

type apiData struct {
	ActionName string
	APIName    string
	Endpoint   string
}

type templateData struct {
	APIs      []apiData
	Namespace string
}

type option struct {
	Mode string `description:"模式" required:"true" short:"m"`
}

// GenerateAPI is 生成api
func GenerateAPI() (ok bool) {
	var opt option
	if _, err := flags.ParseArgs(&opt, os.Args); err != nil {
		ok = false
		return
	}

	ok = true
	wd, _ := os.Getwd()

	var gomod string
	if err := ioos.NewFile(wd, "go.mod").Read(&gomod); err != nil {
		return
	}

	namespaceReg, _ := regexp.Compile(`module\s(.+)`)
	matches := namespaceReg.FindStringSubmatch(gomod)
	data := templateData{
		APIs:      make([]apiData, 0),
		Namespace: matches[1],
	}

	apiReg, _ := regexp.Compile(`func\s(New[a-zA-Z]+API)\(\)`)
	underscore.Chain(
		ioos.NewDirectory(wd, "api").FindDirectories(),
	).Each(func(r ioex.IDirectory, _ int) {
		underscore.Chain(
			r.FindFiles(),
		).Each(func(cr ioex.IFile, _ int) {
			filename := cr.GetName()
			if strings.Contains(filename, "_test") {
				return
			}

			var text string
			if err := cr.Read(&text); err != nil {
				return
			}

			matches := apiReg.FindStringSubmatch(text)
			ext := cr.GetExt()
			data.APIs = append(data.APIs, apiData{
				APIName:    strings.Replace(filename, ext, "", 1),
				ActionName: matches[1],
				Endpoint:   r.GetName(),
			})
		})
	})

	tpl, err := template.New("api-metadata").Parse(`package api

import ({{ range .APIs }}
	"{{ $.Namespace }}/api/{{ .Endpoint }}"{{ end }}
	"github.com/ahl5esoft/lite-go/api"
)

func init() {{ "{" }}{{ range .APIs }}
	api.Register("{{ .Endpoint }}", "{{ .APIName }}", {{ .Endpoint }}.{{ .ActionName }}()){{ end }}
}`)
	if err != nil {
		return
	}

	f, err := ioos.NewFile(wd, "api", "metadata.go").GetFile()
	if err != nil {
		return
	}

	defer f.Close()
	tpl.Execute(f, data)
	return
}
