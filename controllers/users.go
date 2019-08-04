package controllers

import (
	"fmt"
	"net/http"

	"gophr.com/views"
)

// NewUsers parses the templates related to the user and stores them in Users struct
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("base", "views/users/new.gohtml"),
	}
}

// New function is used to render the signup form (for creating a new user)
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "From create function in user controller")
}

// Users will hold processed templates related to user operations
type Users struct {
	NewView *views.View
}
