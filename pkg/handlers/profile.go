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
	tokenStr, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/sign-in/", http.StatusMovedPermanently)
		return
	}

	claims, errValid := utils.ValidateJWT(tokenStr.Value)

	if errValid != nil {
		cookie := &http.Cookie{Path: "/", Name: "token", Value: "", HttpOnly: true, MaxAge: 1}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/sign-in/", http.StatusMovedPermanently)
		return
	}

	user, _ := dal.GetUserById(claims.Id)
  
	// println("user.Id", user.Id)
	// println("user.FirstName", user.FirstName)
	// println("user.LastName", user.LastName)
	// println("user.Birthday", user.Birthday)
	// println("user.City", user.City)
	// println("user.Email", user.Email)
	// println("user.Interests", user.Interests)
	// println("user.Password", user.Password)

	profileTemplate, _ := template.ParseFiles("html/profile.html")
	profileTemplate.Execute(w, user)
}
