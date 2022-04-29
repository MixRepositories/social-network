package handlers

import (
	"highload-architect/pkg/utils"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHome(w, r)
	default:
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	tokenStr, err := r.Cookie("token")

	if err == nil {
		_, errValid := utils.ValidateJWT(tokenStr.Value)

		if errValid == nil {
			http.Redirect(w, r, "/profile/", http.StatusMovedPermanently)
			return
		}
	}

	homeTemplate, _ := template.ParseFiles("html/main.html")
	homeTemplate.Execute(w, "Home page")
}
