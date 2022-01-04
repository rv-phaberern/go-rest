package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Server is where all of our http handlers live
type Server struct {
	store *Store
}

// NewServer ...
func NewServer(store *Store) *Server {
	return &Server{
		store: store,
	}
}

// HandleGetHealth ...
func (s *Server) HandleGetHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	}
}

// HandleGetUsers ...
func (s *Server) HandleGetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.store.GetUsers()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error getting data - %s", err.Error())))
			return
		}

		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			fmt.Println("error marshalling shit", err.Error())
			return
		}
	}
}

// HandleGetUserByID ...
func (s *Server) HandleGetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "userID")

		if idParam == "" {
			w.Write([]byte("missing userID query param"))
			return
		}

		userID, err := strconv.Atoi(idParam)

		user, err := s.store.GetUserByID(userID)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error getting data - %s", err.Error())))
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Println("error marshalling shit", err.Error())
			return
		}
	}
}

// HandleAddUser ...
func (s *Server) HandleAddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error encoding request body - %s", err.Error())))
			return
		}

		err = s.store.AddUser(&user)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error adding user - %s", err.Error())))
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Println("error marshalling shit", err.Error())
			return
		}
	}
}

// HandleUpdateUser ...
func (s *Server) HandleUpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error encoding request body - %s", err.Error())))
			return
		}

		err = s.store.UpdateUser(&user)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error adding user - %s", err.Error())))
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Println("error marshalling shit", err.Error())
			return
		}
	}
}
