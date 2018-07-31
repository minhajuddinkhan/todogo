package store

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/models"
)

//GetTodos GetTodos
func (pg *PgStore) GetTodos(todos *[]models.Todo) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Preload("User").Find(todos)
}

//GetTodoByID GetTodoByID
func (pg *PgStore) GetTodoByID(todo *models.Todo, todoID string) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.First(todo, todoID)

}
