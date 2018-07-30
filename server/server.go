package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

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

	logrus.Info(fmt.Sprintf("Server listening on %s", address))
	err := s.httpServer.ListenAndServe()
	if err != nil {
		logrus.Error(fmt.Sprintf("Server crashed on bootstrap: error => %s", err))
		return
	}

}

func (s *Server) AppendDatabaseToSvr(db *gorm.DB) {
	s.Database = db

}
