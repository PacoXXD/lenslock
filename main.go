package main

import (
	"fmt"
	"net/http"

	"github.com/PacoXXD/lenslock/controller"
	"github.com/PacoXXD/lenslock/templates"
	"github.com/PacoXXD/lenslock/views"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// func execTemplate(w http.ResponseWriter, filepath string) {
// 	t, err := views.Parse(filepath)
// 	if err != nil {
// 		fmt.Printf("Parse html failed %s", err)
// 		http.Error(w, "There is an error parsing the template.", http.StatusInternalServerError)
// 		return
// 	}

// 	// t.Execute(w, nil)
// 	// data, err := os.ReadFile(filepath)
// 	// if err != nil {
// 	// 	fmt.Printf("Read html failed %s", err)
// 	// 	http.Error(w, "There is an error reading the template.", http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	// w.Write(data)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpPath := filepath.Join("templates", "home.gohtml") //lenslocked/templates/contact.gohtml
// 	execTemplate(w, tmpPath)
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpPath := filepath.Join("templates", "contact.gohtml")
// 	execTemplate(w, tmpPath)
// }

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpPath := filepath.Join("templates", "support.gohtml")
// 	execTemplate(w, tmpPath)
// http.ServeFile(w, r, "support.html")

// w.Header().Set("Content-Type", "text/html; charset=utf-8")
// data, err := os.ReadFile("support.html")
// if err != nil {
// 	fmt.Printf("Read html failed %s", err)
// 	return
// }
// w.Write(data)

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}

// }

// type Router struct {
// }

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

func main() {
	// mux := http.NewServeMux()
	// http.HandleFunc("/", pathHandler)
	// http.HandlerFunc

	// var router Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gotmp"))))
	r.Get("/contact", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gotmp"))))
	r.Get("/faq", controller.StaticHandler(views.Must(views.ParseFS(templates.FS, "support.gotmp"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not Found", http.StatusNotFound)

	})

	// err := http.ListenAndServe(":8080", mux)
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", r)
}
