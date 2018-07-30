package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/urfave/negroni"
)

type RouterConf struct {
	Negroni *negroni.Negroni
	Router  *mux.Router
}

//CreateRouter CreateRouter
func CreateRouter() *RouterConf {

	n := &RouterConf{
		Negroni: negroni.Classic(),
		Router:  mux.NewRouter(),
	}
	return n

}

//AppendHandlerFunc appends a hanlder function to negroni hanlder
func (n *RouterConf) AppendHandlerFunc(handlerFunc func(w http.ResponseWriter, r *http.Request)) {

	n.Negroni.UseHandlerFunc(handlerFunc)
}

//AppendHandler appends an http handler
func (n *RouterConf) AppendHandler(handler http.Handler) {
	n.Negroni.UseHandler(handler)
}

func (n *RouterConf) RegisterHandler() {

	n.Negroni.UseHandler(n.Router)
}

func (n *RouterConf) RegisterHandlerFunc(method string, path string, handler http.HandlerFunc) {

	n.Router.HandleFunc(path, handler).Methods(method)
}
