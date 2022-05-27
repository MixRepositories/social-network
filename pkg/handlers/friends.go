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
	case "POST":
		postFriends(w, r)
	case "DELETE":
		deleteFriends(w, r)
	default:
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
	}
}

func getFriends(w http.ResponseWriter, r *http.Request) {
	claims, errClaims := utils.CheckAuthRedirect(w, r)
	if errClaims != nil {
		return
	}

	friends, friendsErr := dal.GetFriends(claims.Id)

	if friendsErr != nil {
		fmt.Println("friendsErr", friendsErr)
		http.Error(w, "Ошибка!", http.StatusInternalServerError)
		return
	}

	var exception []uint16
	for i := 0; i < len(friends); i++ {
		exception = append(exception, friends[i].Id)
	}
	exception = append(exception, claims.Id)

	var tmpData TmpData

	users, usersErr := dal.GetUsers(exception)

	if usersErr != nil {
		fmt.Println(usersErr)
		http.Error(w, "Ошибка!", http.StatusInternalServerError)
		return
	}

	tmpData.Users = users
	tmpData.Friends = friends

	homeTemplate, _ := template.ParseFiles("html/friends.html")
	homeTemplate.Execute(w, tmpData)
}

func postFriends(w http.ResponseWriter, r *http.Request) {
	claims, errClaims := utils.CheckAuthRedirect(w, r)
	if errClaims != nil {
		return
	}

	friendId := r.FormValue("friendId")

	friendsErr := dal.CreateFriends(claims.Id, friendId)

	if friendsErr != nil {
		fmt.Println("friendsErr", friendsErr)
		http.Error(w, "Ошибка!", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/friends/", http.StatusMovedPermanently)
}

func deleteFriends(w http.ResponseWriter, r *http.Request) {
	claims, errClaims := utils.CheckAuthRedirect(w, r)
	if errClaims != nil {
		return
	}

	query := r.URL.Query()
	friendIds, present := query["friendId"]
	if !present || len(friendIds) == 0 {
		http.Error(w, "friendId not present!", http.StatusInternalServerError)
		return
	}

	friendId := friendIds[0]
	selfId := claims.Id

	errDeleteFriend := dal.DeleteFriend(selfId, friendId)
	if errDeleteFriend != nil {
		http.Error(w, "Ошибка!", http.StatusInternalServerError)
		return
	}
}
