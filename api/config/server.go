package config

import "strconv"

// ServerConfig server config interface
type ServerConfig interface {
	Port() string
	Mode() string
}

// Server server config struct
type Server struct {
	port string
	mode string
}

// NewServerConfig create server config struct instance
func NewServerConfig() *Server {
	port := envInt("API_PORT", defaultPort)
	mode := envString("MODE", "debug")

	if mode != "release" && mode != "debug" {
		panic("Unavailable gin mode")
	}

	return &Server{
		port: strconv.Itoa(port),
		mode: mode,
	}
}

// Port get server port number
func (server *Server) Port() string {
	return server.port
}

// Mode get server mode
func (server *Server) Mode() string {
	return server.mode
}