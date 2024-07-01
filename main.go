package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PacoXXD/lenslock/controller"
	"github.com/PacoXXD/lenslock/models"
	"github.com/PacoXXD/lenslock/templates"
	"github.com/PacoXXD/lenslock/views"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "taiwind.gohtml"))))
	r.Get("/contact", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "taiwind.gohtml"))))
	r.Get("/faq", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "support.gohtml", "taiwind.gohtml"))))
	// r.Get("/signup", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "signup.gohtml", "taiwind.gohtml"))))
	r.Get("/signin", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "signin.gohtml", "taiwind.gohtml"))))

	cfg := models.DefaultPostgresConfig()
	pool, err := models.NewPostgresStore(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	userService := models.UserService{
		DB: pool,
	}

	usersC := controller.Users{
		UserService: &userService,
	}

	usersC.Template.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "taiwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not Found", http.StatusNotFound)

	})

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}

}
