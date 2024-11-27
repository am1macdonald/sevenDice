package templates

import (
	"html/template"
	"path/filepath"
)

type TemplateHandler struct {
	templates map[string]*template.Template
}

type PageData struct {
	Title  string
	Groups []Group
}

type Group struct {
	Name        string
	Description string
	PlayerCount int
	MaxPlayers  int
}

//func New() *Templates {
//	return &Templates{
//		templates: template.Must(template.ParseGlob("templates/**/*.html")),
//	}
//}
//
//func (t *Templates) Render(w io.Writer, name string, data interface{}, c context.Context) error {
//	return t.templates.ExecuteTemplate(w, name, data)
//}

func NewTemplateHandler() (*TemplateHandler, error) {
	templates := make(map[string]*template.Template)

	baseLayout := filepath.Join("templates", "layouts", "base.html")

	pages, err := filepath.Glob(filepath.Join("templates", "pages", "*.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.ParseFiles(baseLayout, page)
		if err != nil {
			return nil, err
		}
		templates[name] = tmpl
	}

	return &TemplateHandler{templates: templates}, nil
}

func (th *TemplateHandler) GetTemplate(name string) *template.Template {
	return th.templates[name]
}
