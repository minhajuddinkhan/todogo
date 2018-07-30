package store

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/db"
	"github.com/minhajuddinkhan/todogo/models"
)

//TODO:: MAKE A STORE StRUCt

//GetTodos
func GetTodos(pg *db.PostgresDB, todos *[]models.Todo) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Preload("User").Find(todos)
}

//GetTodoById GetTodoById
func GetTodoById(pg *db.PostgresDB, todo *models.Todo, todoID string) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.First(todo, todoID)

}
