package controllers

import (
	"gophr.com/views"
)

// NewStatic processes templates of static pages and assigns them to a Static type
func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("base", "views/static/home.gohtml"),
		Contact:  views.NewView("base", "views/static/contact.gohtml"),
		Faq:      views.NewView("base", "views/static/faq.gohtml"),
		Error404: views.NewView("base", "views/static/error404.gohtml"),
	}
}

// Static stores parsed templates for static pages
type Static struct {
	Home     *views.View
	Contact  *views.View
	Faq      *views.View
	Error404 *views.View
}
