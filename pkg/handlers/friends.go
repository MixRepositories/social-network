package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"highload-architect/pkg/dal"
	"highload-architect/pkg/structs"
	"highload-architect/pkg/utils"
)

type TmpData struct {
	Users   []structs.User
	Friends []structs.User
}

func Friends(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getFriends(w, r)
	default:
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
	}
}

func getFriends(w http.ResponseWriter, r *http.Request) {
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

	var tmpData TmpData

	users, usersErr := dal.GetUsers()

	if usersErr != nil {
		fmt.Println(usersErr)
		http.Error(w, "Ошибка!", http.StatusInternalServerError)
		return
	}

	friends, friendsErr := dal.GetFriends(claims.Id)
	if friendsErr != nil {
		fmt.Println(friendsErr)
		http.Error(w, "Ошибка!", http.StatusInternalServerError)
		return
	}
	tmpData.Users = users
	tmpData.Friends = friends

	homeTemplate, _ := template.ParseFiles("html/friends.html")
	homeTemplate.Execute(w, tmpData)
}
