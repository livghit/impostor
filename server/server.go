package server

import (
	"log/slog"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/livghit/impostor/utils"
	"gorm.io/gorm"
)

// Maybe for later usage atm not needed
type Server struct {
	Configs ServerConfigs
	Router  *chi.Mux
	DB      *gorm.DB
}
type ServerConfigs struct {
	Port string
}

func New(configs *ServerConfigs) *Server {
	if configs == nil {
		configs = &ServerConfigs{
			Port: "8080",
		}
	}

	db, err := gorm.Open(sqlite.Open(utils.DatabasePath()))
	if err != nil {
		panic(err)
	}

	return &Server{
		Configs: *configs,
		Router:  chi.NewRouter(),
		DB:      db,
	}
}

func (s *Server) Start() {
  slog.Info("Server started", "port", s.Configs.Port)
	if err := http.ListenAndServe(s.Configs.Port, s.Router); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
