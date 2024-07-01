package controller

import (
	"fmt"
	"net/http"

	"github.com/PacoXXD/lenslock/models"
)

type Users struct {
	Template struct {
		New Template
	}
	UserService *models.UserService
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Template.New.Execute(w, data)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %v", user)
}
