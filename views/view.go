package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir has the path to the directory containing the layout files
	LayoutDir = "views/layouts/"

	// TemplateDir has the path to the directory containing the templates
	TemplateDir = "views/"

	// TemplateExt has the extension of the template files
	TemplateExt = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// NewView processes the template provided along with layuout files and stores in View type of the page
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View is a struct that holds the processed template of a static page
type View struct {
	Template *template.Template
	Layout   string
}

// Render executes the parsed template that is passed to it
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
