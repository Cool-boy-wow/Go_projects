package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Name string
	Age  int
	Friends []int
}

type Friend struct {
	UserID int
	FriendID int
}

var users []User
var friends []Friend

func main() {
	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/addfriend", addFriend)
	http.HandleFunc("/showfriends", showFriends)
	http.HandleFunc("/changeuserage",changeUserAge)
	http.HandleFunc("/deleteuser", deleteUser)
	http.ListenAndServe(":8080", nil)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = len(users) + 1
	user.Name = r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Age = age
	users = append(users, user)

	fmt.Fprintf(w, "Пользователь зарегистрирован: %+v", user)
}



func showFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "У пользователя нет друзей", http.StatusNotFound)
		return
	}

	if userID < 1 || userID > len(users) {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Друзья пользователя %d: %v", userID, users[userID].Friends)
}

func addFriend(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var friend Friend
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	friend.UserID, _ = strconv.Atoi(r.FormValue("user_id"))
	friend.FriendID, _ = strconv.Atoi(r.FormValue("friend_id"))

	users[friend.UserID].Friends = append(users[friend.UserID].Friends, friend.FriendID)
	users[friend.FriendID].Friends = append(users[friend.FriendID].Friends, friend.UserID)

	fmt.Fprintf(w, "Друг добавлен: %+v", friend)
}


func changeUserAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	

	if userID < 1 || userID > len(users) {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	age, _ := strconv.Atoi(r.FormValue("age"))
	

	users[userID-1].Age = age

	fmt.Fprintf(w, "Возраст пользователя %d изменен на %d", userID, age)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	if userID < 1 || userID > len(users) {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	users = append(users[:userID-1], users[userID:]...)

	fmt.Fprintf(w, "Пользователь %d удален", userID)
}

