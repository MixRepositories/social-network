package handlers

import (
	"net/http"
	"text/template"

	"highload-architect/pkg/dal"
	"highload-architect/pkg/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProfile(w, r)
	default:
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
	}
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	claims, errClaims := utils.CheckAuthRedirect(w, r)
	if errClaims != nil {
		return
	}

	user, _ := dal.GetUserById(claims.Id)

	profileTemplate, _ := template.ParseFiles("html/profile.html")
	profileTemplate.Execute(w, user)
}
