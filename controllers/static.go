package controllers

import (
	"gophr.com/views"
)

// NewStatic processes templates of static pages and assigns them to a Static type
func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("base", "static/home"),
		Contact:  views.NewView("base", "static/contact"),
		Faq:      views.NewView("base", "static/faq"),
		Error404: views.NewView("base", "static/error404"),
	}
}

// Static stores parsed templates for static pages
type Static struct {
	Home     *views.View
	Contact  *views.View
	Faq      *views.View
	Error404 *views.View
}
