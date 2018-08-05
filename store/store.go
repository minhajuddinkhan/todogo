package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/minhajuddinkhan/todogo/models"
)

type Store interface {
	EstablishConnection() *gorm.DB

	SaveSession(session *models.Session) *gorm.DB
	GetSession(sessionModel *models.Session) *gorm.DB
	UpdateSession(session *models.Session) *gorm.DB
	DeleteSession(session *models.Session) *gorm.DB
	GetSessionByUserID(session *models.Session) *gorm.DB

	GetUserByID(User *models.User, userID string) *gorm.DB
	GetUser(User *models.User) *gorm.DB
	GetUserByNameAndPassword(name string, password string, User *models.User) *gorm.DB
	GetUserByName(name string, User *models.User) *gorm.DB

	GetTodos(todos *[]models.Todo) *gorm.DB
	GetTodoByID(todo *models.Todo, todoID string) *gorm.DB
	GetAllUsers(users *[]models.User) *gorm.DB
	CreateTodo(todo *models.Todo) *gorm.DB
	UpdateTodo(todo *models.Todo) *gorm.DB
}
