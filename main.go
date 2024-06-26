package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmartac/lenslocked_v2/controllers"
	"github.com/jmartac/lenslocked_v2/models"
	"github.com/jmartac/lenslocked_v2/templates"
	"github.com/jmartac/lenslocked_v2/views"
)

func main() {
	r := chi.NewRouter()

	dbConfig := models.DefaultPostgresConfig()
	db, err := models.Open(dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	usersCtrl := controllers.Users{
		UserService: &models.UserService{
			DB: db,
		},
	}
	usersCtrl.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersCtrl.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.Get("/signup", usersCtrl.New)
	r.Get("/signin", usersCtrl.SignIn)
	r.Post("/users", usersCtrl.Create)
	r.Post("/signin", usersCtrl.ProcessSignIn)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found ¯\\_(ツ)_/¯", http.StatusNotFound)
	})

	fmt.Println("Starting the server on http://localhost:3000 ...")
	http.ListenAndServe(":3000", r)
}
