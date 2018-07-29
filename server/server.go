package server

import (
	"net/http"
)

//Server TodoAppServer
type Server struct {
	httpServer *http.Server
}

//NewServer Creates a New server
func NewServer() *Server {
	return &Server{}
}

//Listen Listen on an addr
func (s *Server) Listen(address string, handler http.Handler) {

	s.httpServer = &http.Server{
		Addr:    address,
		Handler: handler,
	}
	s.httpServer.ListenAndServe()
}
