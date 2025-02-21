package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) LoadRoutes() *chi.Mux {
	s.Router.Mount("/", loadWebRoutes())
	s.Router.Mount("/api", loadApiRoutes())

	return s.Router
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

func loadWebRoutes() *chi.Mux {
	web := chi.NewRouter()
	web.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Well well"))
	})

	return web
}
