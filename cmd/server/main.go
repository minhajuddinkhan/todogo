package main

import (
	"fmt"
	"log"

	"github.com/minhajuddinkhan/todogo/constants"
	"github.com/minhajuddinkhan/todogo/middlewares"
	"github.com/minhajuddinkhan/todogo/store"

	"github.com/minhajuddinkhan/todogo/routes"

	"os"

	"github.com/joho/godotenv"
	config "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/db"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/server"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
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

	conf.Db.ConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"+"",
		conf.Db.Host, conf.Db.Port, conf.Db.Username, conf.Db.Name, conf.Db.Password)

	app.Action = func(c *cli.Context) {

		switch c.Args().First() {
		case "db:initiate":

			todoAppDb := db.NewPostgresDB(conf.Db.ConnectionString, conf.Db.Dialect)
			m := models.GetAllModels()
			todoAppDb.Initialize(m)

		case "serve":

			todoAppStore := store.NewPgStore(conf.Db.ConnectionString)
			todoAppSvr := server.NewServer()

			//ROUTER
			R := router.CreateRouter()
			R.Negroni.UseFunc(middlewares.AuthenticateJWT(constants.UserKey, conf.JWTSecret, constants.Authorization))

			routes.RegisterAllRoutes(*R, conf, todoAppStore)
			R.RegisterHandler()
			todoAppSvr.Listen(":"+conf.Port, R.Negroni)

		}

	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
