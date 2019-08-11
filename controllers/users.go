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
		NewView:   views.NewView("base", "users/new"),
		LogInView: views.NewView("base", "users/login"),
		us:        us,
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
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User is", user)
}

// Login will parse the login form and authenticate users
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalide Email address")
		case models.ErrInvalidPassword:
			fmt.Fprintln(w, "Invalid Password Provided")
		case nil:
			fmt.Fprintln(w, user)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, user)
}

// CookieTest is used to check whether cookie is set or not
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Email is : ", cookie.Value)
}

// Users will hold processed templates related to user operations
type Users struct {
	NewView   *views.View
	LogInView *views.View
	us        *models.UserService
}

// SignupForm contains the details entered by the user in the signup form
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// LoginForm contains the details entered by the user in the login form
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
