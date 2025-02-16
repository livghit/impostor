package server
// Maybe for later usage atm not needed 
type Server struct {
	Configs ServerConfigs
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
	}
}
