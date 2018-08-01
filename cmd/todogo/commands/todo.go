package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/store"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

//Todos Todos
func Todos(store store.Store) *cli.Command {

	return &cli.Command{
		Name:    "todos",
		Aliases: []string{"gotodo todo"},
		Usage:   "Handles Your todos",
		Action: func(c *cli.Context) error {
			color.New(color.FgCyan).Println("What do you wanna do with todos man?")
			cli.ShowAppHelp(c)
			return nil
		},
		Subcommands: []cli.Command{
			*getTodos(store),
		},
	}

}

func getTodos(store store.Store) *cli.Command {

	return &cli.Command{
		Name:    "get",
		Aliases: []string{"gotodo todo get"},
		Usage:   "Fetches All Your todos",
		Action: func(c *cli.Context) error {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Todo", "User", "Priority"})

			windowID := os.Getenv("WINDOWID")
			session := models.Session{
				Session: windowID,
			}

			err := store.GetSession(&session).Error
			if err != nil {

				if gorm.IsRecordNotFoundError(err) {
					fmt.Println("Please log in first.")
					return nil
				}
				panic(err)

			}

			todos := []models.Todo{}

			err = store.GetTodos(&todos).Error
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

			return nil
		},
	}
}
