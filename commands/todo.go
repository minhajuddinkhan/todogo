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

func authenticateSession(store store.Store) (models.Session, error) {

	windowID := os.Getenv("WINDOWID")
	session := models.Session{
		Session: windowID,
	}

	err := store.GetSession(&session).Error
	if err != nil {
		return session, err
	}

	return session, nil

}

//Todos Todos
func Todos(store store.Store) *cli.Command {

	return &cli.Command{
		Name:    "todos",
		Aliases: []string{"todogo todos"},
		Usage:   "Handles Your todos",
		Action: func(c *cli.Context) error {
			color.New(color.FgCyan).Println("What do you wanna do with todos man?")
			cli.ShowAppHelp(c)
			return nil
		},
		Subcommands: []cli.Command{
			*getTodos(store),
			*createTodo(store),
			*updateTodo(store),
		},
	}

}

func updateTodo(store store.Store) *cli.Command {

	return &cli.Command{
		Name:    "update",
		Aliases: []string{"todogo todos update"},
		Usage:   "Updates todos",
		Action: func(c *cli.Context) error {

			session, err := authenticateSession(store)
			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					fmt.Println("Please log in first.")
					return nil
				}
				logrus.Error(err.Error())
				return nil
			}

			fmt.Println("Enter TodoID")
			todo := models.Todo{}
			fmt.Scan(&todo.ID)

			err = store.GetTodoByID(&todo, fmt.Sprint(todo.ID)).Error
			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					fmt.Println("No Todo with ID", todo.ID)
					return nil
				}
			}
			fmt.Println("Enter Todo Name")
			fmt.Scan(&todo.Name)

			fmt.Println("Enter Todo Priority")
			fmt.Scan(&todo.Priorty)
			todo.UserID = session.UserID

			result := store.UpdateTodo(&todo)
			if result.RowsAffected == 0 {
				fmt.Println("No Todos were updated.")
			}
			return nil
		},
	}
}

func getTodos(store store.Store) *cli.Command {

	return &cli.Command{
		Name:    "get",
		Aliases: []string{"todogo todos get"},
		Usage:   "Fetches All Your todos",
		Action: func(c *cli.Context) error {

			_, err := authenticateSession(store)
			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					fmt.Println("Please log in first.")
					return nil
				}
				logrus.Error(err.Error())
				return err

			}

			todos := []models.Todo{}

			err = store.GetTodos(&todos).Error
			if gorm.IsRecordNotFoundError(err) {
				logrus.Error("No todos right now")
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"TodoID", "Todo", "User", "Priority"})
			for i := 0; i < len(todos); i++ {
				table.Append([]string{
					fmt.Sprint(todos[i].ID),
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

func createTodo(store store.Store) *cli.Command {

	return &cli.Command{
		Name:    "create",
		Aliases: []string{"todogo todos create"},
		Usage:   "Creates a todo for you.",
		Action: func(c *cli.Context) error {

			session, err := authenticateSession(store)
			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					fmt.Println("Please log in first.")
					return nil
				}
				logrus.Error(err.Error())
			}
			todo := models.Todo{}
			fmt.Println("Enter Todo Name")
			fmt.Scan(&todo.Name)

			fmt.Println("Enter Todo Priority")
			fmt.Scan(&todo.Priorty)

			todo.UserID = session.UserID
			result := store.CreateTodo(&todo)
			if result.RowsAffected == 0 {
				logrus.Error("Something went wrong", result.Error.Error())
			}
			return nil
		},
	}
}
