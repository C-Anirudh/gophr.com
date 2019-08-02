package main

import (
	"log"
	"net/http"

	"gophr.com/views"

	"github.com/gorilla/mux"
)

var homeView, contactView, faqView, error404View *views.View

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := error404View.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err404 = http.HandlerFunc(error404)
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	faqView = views.NewView("views/faq.gohtml")
	error404View = views.NewView("views/error404.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = err404
	log.Fatal(http.ListenAndServe(":8080", r))
}
