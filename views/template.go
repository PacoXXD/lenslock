package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	// tpl := template.New(patterns[0])
	// tpl = tpl.Funcs(
	// 	template.FuncMap{
	// 		"csrFiled": func() template.HTML {
	// 			return `<-- TODO:Implementthi -->`

	// 		},
	// 	},
	// )
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("Parse FS failed %s", err)
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}

func Parse(filePath string) (Template, error) {

	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		return Template{}, fmt.Errorf("Parse html failed %s", err)
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}

type Template struct {
	HTMLTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.HTMLTpl.Execute(w, nil)
	if err != nil {
		log.Printf("Execute html failed %s", err)
		http.Error(w, "There is an error Execute he template.", http.StatusInternalServerError)
		return
	}
}
