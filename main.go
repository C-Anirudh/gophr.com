package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate, contactTemplate, faqTemplate, error404Template *template.Template

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := error404Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	var err404 = http.HandlerFunc(error404)
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}
	faqTemplate, err = template.ParseFiles("views/faq.gohtml")
	if err != nil {
		panic(err)
	}
	error404Template, err = template.ParseFiles("views/error404.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = err404
	log.Fatal(http.ListenAndServe(":8080", r))
}
