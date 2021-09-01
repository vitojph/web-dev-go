package main

import (
	"fmt"
	"net/http"

	"web-dev-go/views"

	"github.com/gorilla/mux"
)

var (
	homeView     *views.View
	contactView  *views.View
	notFoundView *views.View
	faqView      *views.View
)

// A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	must(contactView.Render(w, nil))
}

func notFound404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	must(notFoundView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	must(faqView.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	notFoundView = views.NewView("bootstrap", "views/404.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")

	var handler404 http.Handler = http.HandlerFunc(notFound404)

	r := mux.NewRouter()
	r.NotFoundHandler = handler404
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	fmt.Println("Serving at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
