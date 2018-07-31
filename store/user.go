package store

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/models"
)

//GetUserByID GetUserByID
func (pg *PgStore) GetUserByID(User *models.User, userID string) *gorm.DB {

	conn := pg.EstablishConnection()
	return conn.First(User, userID)
}

//GetUser GetUser
func (pg *PgStore) GetUser(User *models.User) *gorm.DB {
	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Where(User).First(User)
}
