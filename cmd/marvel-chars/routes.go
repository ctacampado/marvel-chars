package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// SetRoutes is the Router init function
func SetRoutes(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Route("/characters", func(r chi.Router) {
		r.Get("/", GetCharacters)
		r.Get("/{id}", GetCharacterByID)
	})
}
