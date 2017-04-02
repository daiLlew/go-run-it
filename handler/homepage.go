package handler

import (
	"github.com/daiLlew/go-run-it/webModel"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type TemplateHandler struct {
	once      sync.Once
	Filename  string
	templ     *template.Template
	Workspace *webModel.WorkSpace
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.Filename)))
	})
	t.templ.Execute(w, t.Workspace)
}
