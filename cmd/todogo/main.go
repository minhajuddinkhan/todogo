package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minhajuddinkhan/todogo/store"

	"github.com/minhajuddinkhan/todogo/commands"
	"github.com/minhajuddinkhan/todogo/db"

	"github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/constants"
	"github.com/minhajuddinkhan/todogo/middlewares"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/pgstore"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/routes"
	"github.com/minhajuddinkhan/todogo/server"
	"github.com/minhajuddinkhan/todogo/sqlitestore"
	"github.com/urfave/cli"
)

func main() {
	godotenv.Load()
	todoApp := cli.NewApp()
	todoApp.Name = "TODOGO - just a dummy todo application. nothing fancy"

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
	var todoAppStore store.Store

	if conf.Db.Dialect == "postgres" {
		todoAppDb = db.NewPostgresDB(conf)
		todoAppStore = pgstore.NewPgStore(todoAppDb)

	} else {
		conf.Db.Dialect = "sqlite"
		conf.Db.VolumePath = "/tmp/todo.db"
		todoAppDb = db.NewSqliteDB(conf)
		todoAppStore = sqlitestore.NewSqliteStore(todoAppDb)
	}

	todoApp.Commands = []cli.Command{
		*commands.Todos(todoAppStore),
		*commands.Login(todoAppStore),
		*commands.Logout(todoAppStore),
		{
			Name:    "serve",
			Aliases: []string{"todogo serve"},
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
			Aliases: []string{"todogo db"},
			Usage:   "Initiates DB from your command line",
			Action: func(cli *cli.Context) error {
				m := models.GetAllModels()
				todoAppDb.Initialize(m)
				return nil
			},
		},
	}

	err := todoApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
