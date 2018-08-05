package sqlitestore

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/models"
)

//GetUserByID GetUserByID
func (pg *SqliteStore) GetUserByID(User *models.User, userID string) *gorm.DB {

	conn := pg.EstablishConnection()
	return conn.First(User, userID)
}

//GetUser GetUser
func (pg *SqliteStore) GetUser(User *models.User) *gorm.DB {
	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Where(User).First(User)
}

//GetUserByNameAndPassword GetUserByNameAndPassword
func (pg *SqliteStore) GetUserByNameAndPassword(name string, password string, User *models.User) *gorm.DB {
	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Where("name = ? AND password = ?", name, password).First(User)
}

//GetUserByName GetUserByName
func (pg *SqliteStore) GetUserByName(name string, User *models.User) *gorm.DB {
	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Where("name = ?", name).First(User)
}

//GetAllUsers GetAllUsers
func (pg *SqliteStore) GetAllUsers(users *[]models.User) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Find(users)
}

func (pg *SqliteStore) CreateUser(user *models.User) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Create(user)

}
