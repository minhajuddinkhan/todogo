package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/minhajuddinkhan/todogo/store"

	"github.com/jinzhu/gorm"
	"github.com/olekukonko/tablewriter"

	"github.com/joho/godotenv"
	"github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/urfave/cli"
)

func main() {

	//	sessions := map[int]bool{}
	todoCli := cli.NewApp()
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

	todoAppStore := store.NewPgStore(conf.Db.ConnectionString)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Todo", "User", "Priority"})
	todoCli.Action = func(c *cli.Context) {

		switch c.Args().First() {

		case "get":

			windowID := os.Getenv("WINDOWID")
			session := models.Session{
				Session: windowID,
			}

			err := todoAppStore.GetSession(&session).Error
			if err != nil {

				if gorm.IsRecordNotFoundError(err) {
					fmt.Println("Please log in first.")
					return
				}
				panic(err)

			}

			todos := []models.Todo{}

			err = todoAppStore.GetTodos(&todos).Error
			if gorm.IsRecordNotFoundError(err) {
				logrus.Error("No todos right now")
			}

			for i := 0; i < len(todos); i++ {
				table.Append([]string{
					todos[i].Name,
					todos[i].User.Name,
					strconv.Itoa(todos[i].Priorty),
				})
			}

			table.Render()

		case "logout":
			windowID := os.Getenv("WINDOWID")
			fmt.Println("WINDWO ID", windowID)

			session := models.Session{
				Session: windowID,
			}
			err := todoAppStore.DeleteSession(&session).Error
			if err != nil {
				panic(err)
			}
			logrus.Info("Logged out.")
		case "login":

			type LoginRequest struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			c := color.New(color.FgHiGreen)
			rc := color.New(color.FgHiRed)

			var username, password string
			c.Println("Enter Name")
			fmt.Scan(&username)
			c.Println("Enter password")
			fmt.Scan(&password)

			user := models.User{
				Name:     username,
				Password: password,
			}
			err := todoAppStore.GetUser(&user).Error

			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					rc.Println("Bad credentials")
					return
				}
				rc.Println("Something bad happened\n" + err.Error())
				return
			}

			windowID := (os.Getenv("WINDOWID"))
			fmt.Println("WINDWO ID", windowID)

			userSession := models.Session{
				UserID: user.ID,
			}

			if todoAppStore.GetSessionByUserID(&userSession).RowsAffected == 0 {
				err = todoAppStore.SaveSession(&models.Session{
					Session: windowID,
					UserID:  user.ID,
				}).Error

				if err != nil {
					panic(err)
				}
				return
			}
			err = todoAppStore.UpdateSession(&models.Session{
				Session: windowID,
				UserID:  user.ID,
			}).Error

			if err != nil {
				panic(err)
			}
		case "todos":
			todos := []models.Todo{}

			err = todoAppStore.GetTodos(&todos).Error
			if gorm.IsRecordNotFoundError(err) {
				fmt.Println("No todos right now")
			}

			for i := 0; i < len(todos); i++ {
				table.Append([]string{
					todos[i].Name,
					todos[i].User.Name,
					strconv.Itoa(todos[i].Priorty),
				})
			}

			table.Render()

		}
	}

	err = todoCli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
