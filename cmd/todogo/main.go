package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
	config "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/constants"
	"github.com/minhajuddinkhan/todogo/db"
	"github.com/minhajuddinkhan/todogo/middlewares"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/server"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf := &config.Configuration{
		JWTSecret: os.Getenv("JWTSECRET"),
		Port:      os.Getenv("SVR_PORT"),
		Db: config.Db{
			Dialect:  "postgres",
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s"+"",
		conf.Db.Host, conf.Db.Port, conf.Db.Username, conf.Db.Name, conf.Db.Password)

	fmt.Println(connectionString)
	//SERVER
	todoAppSvr := server.NewServer()
	todoAppDb := db.NewPostgresDB(connectionString, conf.Db.Dialect)
	conn := todoAppDb.EstablishConnection()
	todoAppSvr.AppendDatabaseToSvr(conn)
	m := models.GetAllModels()
	todoAppDb.Migrate(m)
	// //ROUTER
	R := router.CreateRouter()
	R.Negroni.UseFunc(middlewares.AppendDatabaseContext(constants.DbKey, todoAppSvr.Database))
	R.Negroni.UseFunc(middlewares.AuthenticateJWT(constants.UserKey, conf.JWTSecret, constants.Authorization))

	R.Router.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "YOLO!")
	})

	R.RegisterHandler()
	todoAppSvr.Listen(":"+conf.Port, R.Negroni)
}
