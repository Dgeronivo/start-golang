package templateprovider

import (
	"os"
	"text/template"
)

func Provide(tmplName string) (*template.Template, error) {
	wd, _ := os.Getwd()

	return template.ParseFiles(wd + "\\web\\template\\" + tmplName)
}
