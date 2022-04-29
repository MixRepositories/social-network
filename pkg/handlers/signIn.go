package handlers

import (
	"net/http"
	"text/template"

	"highload-architect/pkg/dal"
	"highload-architect/pkg/utils"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getSignIn(w, r)
	case "POST":
		postSignIn(w, r)
	default:
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
	}
}

func getSignIn(w http.ResponseWriter, r *http.Request) {
	signInTemplate, _ := template.ParseFiles("html/sign-in.html")
	signInTemplate.Execute(w, "Sing in page")
}

func postSignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	hashPassword := utils.GetHashString(password)

	user, userErr := dal.GetAuthUser(email, hashPassword)
	if userErr != nil {
		http.Error(w, "Пользователь не найден!", http.StatusNotFound)
		return
	}

	if user.Email != email {
		http.Error(w, "Пользователь не найден!", http.StatusNotFound)
		return
	}

	token, jwtErr := utils.GenerateJWT(user.Id, user.Email, user.FirstName, user.LastName)
	if jwtErr != nil {
		http.Error(w, "Ошибка при генерации токена!", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{Path: "/", Name: "token", Value: token, HttpOnly: true, MaxAge: 60 * 60 * 24}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/profile/", http.StatusMovedPermanently)
}
