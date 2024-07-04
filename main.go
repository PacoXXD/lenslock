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
	"github.com/gorilla/csrf"
)

func main() {
	fmt.Println("Starting the server setup...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fmt.Println("Setting up routes...")

	r.Get("/", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "taiwind.gohtml"))))
	r.Get("/contact", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "taiwind.gohtml"))))
	r.Get("/faq", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "support.gohtml", "taiwind.gohtml"))))
	r.Get("/signin", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "signin.gohtml", "taiwind.gohtml"))))

	fmt.Println("Routes set up. Connecting to the database...")

	cfg := models.DefaultPostgresConfig()
	pool, err := models.NewPostgresStore(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	fmt.Println("Database connected. Setting up user service...")

	userService := models.UserService{
		DB: pool,
	}
	sessionService := models.SessionService{
		DB: pool,
	}

	usersC := controller.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}

	usersC.Template.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "taiwind.gohtml"))
	usersC.Template.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "taiwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/user/me", usersC.CurrentUser)
	r.Post("/signout", usersC.ProcessSignOut)

	fmt.Println("User service set up. Configuring 404 handler...")

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not Found", http.StatusNotFound)
	})
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect([]byte(csrfKey), csrf.Secure(false))

	fmt.Println("Starting the HTTP server on port 3000...")
	err = http.ListenAndServe(":3000", csrfMw(r))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
}
