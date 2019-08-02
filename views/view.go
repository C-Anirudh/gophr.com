package views

import (
	"html/template"
)

// NewView processes the template provided along with layuout files and stores in View type of the page
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// View is a struct that holds the processed template of a static page
type View struct {
	Template *template.Template
}
