package sqlitestore

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/models"
)

//GetTodos GetTodos
func (pg *SqliteStore) GetTodos(todos *[]models.Todo) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Preload("User").Find(todos)
}

//GetTodoByID GetTodoByID
func (pg *SqliteStore) GetTodoByID(todo *models.Todo, todoID string) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.First(todo, todoID)

}
func (pg *SqliteStore) CreateTodo(todo *models.Todo) *gorm.DB {
	conn := pg.EstablishConnection()
	defer conn.Close()

	return conn.Create(todo)
}

func (pg *SqliteStore) UpdateTodo(todo *models.Todo) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Save(todo)
}
