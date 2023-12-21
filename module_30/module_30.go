package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
	Id      int      `json:"id"`
}

func (u *User) toString() string {
	return fmt.Sprintf("Name is %s and age is %d and friends are %s and ID is %d \n", u.Name, u.Age, strings.Join(u.Friends, ","), u.Id)
}

type service struct {
	store map[int]*User
}

var new_id int

func main() {
	mux := http.NewServeMux()
	srv := service{make(map[int]*User)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.Get)
	mux.HandleFunc("/makefriends", srv.MakeFriends)
	http.ListenAndServe("localhost:8080", mux)
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

		}
		defer r.Body.Close()
		var u User
		new_id += 1
		u.Id = new_id
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		s.store[u.Id] = &u
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + u.Name))
		return

	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, user := range s.store {
			response += user.toString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

		}
		defer r.Body.Close()
		var m []byte
		if err := json.Unmarshal(content, &m); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(m))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
