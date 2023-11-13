package main

import (
	"net/http"

	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

var users = []User{
	{ID: "1", Name: "John Doe", Bio: "This is John Doe"},
	{ID: "2", Name: "Some other", Bio: "Some other bio"},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	r.Post("/user", addUser)

	http.ListenAndServe(":3333", r)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.ContentLength == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == "" || user.Name == "" || user.Bio == "" {
		http.Error(w, "User fields cannot be empty", http.StatusBadRequest)
		return
	}

	users = append(users, user)

	// Return updated user list as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}