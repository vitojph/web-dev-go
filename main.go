package main

import (
	"fmt"
	"net/http"

	"web-dev-go/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var notFoundView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}

func notFound404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := notFoundView.Template.ExecuteTemplate(w, notFoundView.Layout, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	notFoundView = views.NewView("bootstrap", "views/404.gohtml")

	var handler404 http.Handler = http.HandlerFunc(notFound404)

	r := mux.NewRouter()
	r.NotFoundHandler = handler404
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	fmt.Println("Serving at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
