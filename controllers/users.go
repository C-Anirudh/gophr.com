package controllers

import (
	"fmt"
	"net/http"

	"gophr.com/models"
	"gophr.com/views"
)

// NewUsers parses the templates related to the user and stores them in Users struct
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("base", "users/new"),
		us:      us,
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
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User is", user)
}

// Users will hold processed templates related to user operations
type Users struct {
	NewView *views.View
	us      *models.UserService
}

// SignupForm contains the details entered by the user in the signup form
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
