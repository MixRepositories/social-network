package handlers

import (
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getSignOut(w, r)
	default:
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
	}
}

func getSignOut(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{Path: "/", Name: "token", Value: "", HttpOnly: true, MaxAge: 1}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/sign-in/", http.StatusMovedPermanently)
}
