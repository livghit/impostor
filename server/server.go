package server

import "net/http"

// Maybe for later usage atm not needed
type Server struct {
	Configs ServerConfigs
	Router  http.Handler
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

	return &Server{
		Configs: *configs,
		Router:  loadAllRoutes(),
	}
}
