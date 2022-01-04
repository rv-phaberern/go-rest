package main

import (
	"fmt"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func configureRouter(server *Server) *chi.Mux {
	fmt.Println("configuring router")
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", server.HandleGetHealth())
	r.Get("/users", server.HandleGetUsers())
	r.Get("/users/{userID}", server.HandleGetUserByID())
	r.Post("/users", server.HandleAddUser())
	r.Put("/users", server.HandleUpdateUser())

	return r
}
