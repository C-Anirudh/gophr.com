package main

import (
	"log"
	"net/http"

	"gophr.com/views"

	"github.com/gorilla/mux"
)

var homeView, contactView, faqView, error404View, signupView *views.View

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	must(error404View.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err404 = http.HandlerFunc(error404)
	homeView = views.NewView("base", "views/home.gohtml")
	contactView = views.NewView("base", "views/contact.gohtml")
	faqView = views.NewView("base", "views/faq.gohtml")
	error404View = views.NewView("base", "views/error404.gohtml")
	signupView = views.NewView("base", "views/signup.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/signup", signup)
	r.NotFoundHandler = err404
	log.Fatal(http.ListenAndServe(":8080", r))
}
