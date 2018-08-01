package main

import (
	"fmt"
	"log"
	"os"

	"github.com/minhajuddinkhan/todogo/cmd/todogo/commands"

	"github.com/joho/godotenv"

	"github.com/minhajuddinkhan/todogo/cmd/todogo/routes"
	"github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/constants"
	"github.com/minhajuddinkhan/todogo/db"
	"github.com/minhajuddinkhan/todogo/middlewares"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/server"
	"github.com/minhajuddinkhan/todogo/store"
	"github.com/urfave/cli"
)

func main() {
	todoApp := cli.NewApp()
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
	todoAppDb := db.NewPostgresDB(conf.Db.ConnectionString)
	todoAppStore := store.NewPgStore(todoAppDb)

	todoApp.Commands = []cli.Command{
		*commands.Todos(todoAppStore),
		*commands.Login(todoAppStore),
		*commands.Logout(todoAppStore),
		{
			Name:    "serve",
			Aliases: []string{"gotodo serve"},
			Usage:   "Starts serve from your command line",
			Action: func(cli *cli.Context) error {
				todoAppSvr := server.NewServer()

				//ROUTER
				R := router.CreateRouter()
				R.Negroni.UseFunc(middlewares.AuthenticateJWT(constants.UserKey, conf.JWTSecret, constants.Authorization))

				routes.RegisterAllRoutes(*R, conf, todoAppStore)
				R.RegisterHandler()
				todoAppSvr.Listen(":"+conf.Port, R.Negroni)
				return nil
			},
		},
		{
			Name:    "db",
			Aliases: []string{"gotodo db"},
			Usage:   "Initiates DB from your command line",
			Action: func(cli *cli.Context) error {
				m := models.GetAllModels()
				todoAppDb.Initialize(m)
				return nil
			},
		},
	}

	err = todoApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
