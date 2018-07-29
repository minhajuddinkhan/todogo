package main

import (
	"fmt"
	"net/http"

	"github.com/minhajuddinkhan/todogo/middlewares"
	router "github.com/minhajuddinkhan/todogo/router"
	server "github.com/minhajuddinkhan/todogo/server"
)

type TodoAppConfig struct {
	JwtSecret string
}

const (
	UserKey       = iota
	Authorization = "Authorization"
)

func main() {

	conf := &TodoAppConfig{
		JwtSecret: "ILoveGoLang!",
	}

	todoAppSvr := server.NewServer()
	R := router.CreateRouter()

	R.Negroni.UseFunc(middlewares.AuthenticateJWT(UserKey, conf.JwtSecret, Authorization))
	R.Router.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "YOLO!")

	})

	R.RegisterHandler()
	todoAppSvr.Listen(":3000", R.Negroni)
}
