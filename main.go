package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmartac/lenslocked_v2/controllers"
	"github.com/jmartac/lenslocked_v2/templates"
	"github.com/jmartac/lenslocked_v2/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found ¯\\_(ツ)_/¯", http.StatusNotFound)
	})

	fmt.Println("Starting the server on http://localhost:3000 ...")
	http.ListenAndServe(":3000", r)
}
