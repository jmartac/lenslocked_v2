package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmartac/lenslocked_v2/controllers"
	"github.com/jmartac/lenslocked_v2/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.Parse("templates/home.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.Parse("templates/contact.gohtml"))))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.Parse("templates/faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found ¯\\_(ツ)_/¯", http.StatusNotFound)
	})

	fmt.Println("Starting the server on http://localhost:3000 ...")
	http.ListenAndServe(":3000", r)
}
