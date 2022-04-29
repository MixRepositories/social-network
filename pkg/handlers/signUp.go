package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"highload-architect/pkg/dal"
	"highload-architect/pkg/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getSignUp(w, r)
	case "POST":
		postSignUp(w, r)
	default:
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
	}

}

func getSignUp(w http.ResponseWriter, r *http.Request) {
	signUpTemplate, _ := template.ParseFiles("html/sign-up.html")
	signUpTemplate.Execute(w, "Sing up page")
}

func postSignUp(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	gender := r.FormValue("gender")
	age := r.FormValue("age")
	city := r.FormValue("city")
	interests := r.FormValue("interests")
	birthday := r.FormValue("birthday")

	hashPassword := utils.GetHashString(password)

	createErr := dal.CreateUser(
		email,
		hashPassword,
		firstName,
		lastName,
		birthday,
		gender,
		age,
		city,
		interests,
	)

	if createErr != nil {
		fmt.Println(createErr)
		http.Error(w, "Ошибка при создание пользователя!", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/sign-in/", http.StatusMovedPermanently)
}
