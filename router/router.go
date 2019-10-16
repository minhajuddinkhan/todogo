package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Route struct {
	Method  string
	URI     string
	Handler http.HandlerFunc
}

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

//RegisterMultipleHandlers RegisterMultipleHandlers
func (n *RouterConf) RegisterMultipleHandlers(routes []Route) {

	for _, r := range routes {
		n.Router.HandleFunc(r.URI, r.Handler).Methods(r.Method)
	}
}
