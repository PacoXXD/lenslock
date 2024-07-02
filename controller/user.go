package controller

import (
	"fmt"
	"net/http"

	"github.com/PacoXXD/lenslock/models"
)

type Users struct {
	Template struct {
		New    Template
		SignIn Template
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
		// CSRFFild template.HTML
	}
	data.Email = r.FormValue("email")
	// data.CSRFFild = csrf.TemplateField(r)
	u.Template.New.Execute(w, data)
}

func (u *Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Template.SignIn.Execute(w, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Wrong password or email", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
	setCookie(w, CookieSession, session.TokenHash)
	http.Redirect(w, r, "/user/me", http.StatusFound)
	fmt.Fprintf(w, "User authenticated: %v", user)

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
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
	setCookie(w, CookieSession, session.TokenHash)
	http.Redirect(w, r, "/user/me", http.StatusFound)

	fmt.Fprintf(w, "User created: %v", user)
}

// func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
// 	tokenCookie, err := readCookie(r, CookieSession)
// 	if err != nil {
// 		fmt.Fprint(w, "No cookie")
// 		http.Redirect(w, r, "/signin", http.StatusFound)
// 		return
// 	}
// 	fmt.Println(tokenCookie)
// 	user, err := u.SessionService.User(tokenCookie)
// 	fmt.Println(user)
// 	if err != nil {
// 		fmt.Fprint(w, "Invalid token")
// 		http.Redirect(w, r, "/signin", http.StatusFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Current user: %v\n", user)
// 	fmt.Fprintf(w, "header: %v", r.Header)

// }

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := readCookie(r, "CookieSession")
	if err != nil {
		fmt.Fprint(w, "No cookie")
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	fmt.Println(tokenCookie)
	user, err := u.SessionService.User(tokenCookie)
	if err != nil {
		fmt.Fprint(w, "Invalid token")
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "Current user: %v\n", user)
	fmt.Fprintf(w, "header: %v", r.Header)
}

// func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
// 	token, err := readCookie(r, CookieSession)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Redirect(w, r, "/signin", http.StatusFound)
// 		return
// 	}
// 	user, err := u.SessionService.User(token)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Redirect(w, r, "/signin", http.StatusFound)
// 		return
// 	}
// 	fmt.Fprintf(w, "Current user: %s\n", user.Email)
// }
