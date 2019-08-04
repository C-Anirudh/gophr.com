package controllers

import (
	"fmt"
	"net/http"

	"gophr.com/views"
)

// NewUsers parses the templates related to the user and stores them in Users struct
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("base", "users/new"),
	}
}

// New function is used to render the signup form (for creating a new user)
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create will parse the sign up form and create a new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}

// Users will hold processed templates related to user operations
type Users struct {
	NewView *views.View
}

// SignupForm contains the details entered by the user in the signup form
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
