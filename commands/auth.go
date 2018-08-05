package commands

import (
	"fmt"
	"os"

	"github.com/minhajuddinkhan/todogo/store"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/models"
)

//Login Login
func Login(store store.Store) *cli.Command {

	return &cli.Command{
		Name:    "login",
		Aliases: []string{"todogo login"},
		Usage:   "Starts session from your command line",
		Action: func(cli *cli.Context) error {

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
			err := store.GetUser(&user).Error

			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					rc.Println("Bad credentials")
					return nil
				}
				rc.Println("Something bad happened\n" + err.Error())
				return nil
			}

			windowID := (os.Getenv("WINDOWID"))

			userSession := models.Session{
				UserID: user.ID,
			}

			if store.GetSessionByUserID(&userSession).RowsAffected == 0 {
				err = store.SaveSession(&models.Session{
					Session: windowID,
					UserID:  user.ID,
				}).Error

				if err != nil {
					rc.Println("Something bad happened\n" + err.Error())
					return nil
				}
			}
			err = store.UpdateSession(&models.Session{
				Session: windowID,
				UserID:  user.ID,
			}).Error

			if err != nil {
				rc.Println("Something bad happened\n" + err.Error())
			}
			logrus.Info("Login successfull!")
			return nil
		},
	}
}

//Logout Logout
func Logout(store store.Store) *cli.Command {
	return &cli.Command{
		Name:    "logout",
		Aliases: []string{"todogo logout"},
		Usage:   "Ends session from your command line",
		Action: func(c *cli.Context) error {
			windowID := os.Getenv("WINDOWID")
			session := models.Session{
				Session: windowID,
			}
			err := store.DeleteSession(&session).Error
			if err != nil {
				return err
			}
			logrus.Info("Logged out.")
			return nil
		},
	}
}
