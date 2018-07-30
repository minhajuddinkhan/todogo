package server

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

//Server TodoAppServer
type Server struct {
	httpServer *http.Server
	Database   *gorm.DB
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

func (s *Server) AppendDatabaseToSvr(db *gorm.DB) {
	s.Database = db

}
