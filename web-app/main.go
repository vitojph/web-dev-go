package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>¡Bienvenido a mi web!</h1>")
	fmt.Fprint(w, "<p>Es mía, sí.</p>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contacto</h1>")
	fmt.Fprint(w, "<p>Escríbeme un email a <a href=\"mailto:contact@mail.com\">esta dirección</a>.</p>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Preguntas Frecuentes</h1>")
	fmt.Fprint(w, "<p>Esta es mi lista de preguntas frecuentes.</p>")
}

func notFound404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 Personalizada</h1>")
	fmt.Fprint(w, "<p>Lo siento, página no encontrada.</p>")
}

func main() {
	var handler404 http.Handler = http.HandlerFunc(notFound404)
	r := mux.NewRouter()
	r.NotFoundHandler = handler404
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":8080", r)
}
