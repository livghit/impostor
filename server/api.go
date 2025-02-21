package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadAllRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Well well"))
	})

	r.Mount("/api", loadApiRoutes())

  return r
}

func loadApiRoutes() *chi.Mux {
	api := chi.NewRouter() // Route responsible for creating a player
	api.Post("/create-player", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We create a player now"))
	})
	// Route responsible for creating a lobby
	api.Post("/create-lobby", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We create a lobby now"))
	})
	// Route responsible for creating a game
	api.Post("/create-game", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We create a game now"))
	})

	return api
}
