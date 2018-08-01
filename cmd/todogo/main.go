package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/minhajuddinkhan/todogo/cmd/todogo/commands"
	"github.com/minhajuddinkhan/todogo/db"

	"github.com/joho/godotenv"

	"github.com/minhajuddinkhan/todogo/cmd/todogo/routes"
	"github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/constants"
	"github.com/minhajuddinkhan/todogo/middlewares"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/pgstore"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/server"
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
			Dialect:    os.Getenv("DB_DIALECT"),
			Host:       os.Getenv("DB_HOST"),
			Port:       os.Getenv("DB_PORT"),
			Name:       os.Getenv("DB_NAME"),
			Username:   os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			VolumePath: os.Getenv("DB_VOULME_PATH"),
		},
	}

	var todoAppDb db.Database

	fmt.Println(conf.Db.Dialect)
	if conf.Db.Dialect == "postgres" {
		todoAppDb = db.NewPostgresDB(conf)

	} else if conf.Db.Dialect == "sqlite" {
		todoAppDb = db.NewSqliteDB(conf)

	} else {
		logrus.Error("Invalid Dialect" + conf.Db.Dialect)
		return
	}

	todoAppStore := pgstore.NewPgStore(todoAppDb)

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
