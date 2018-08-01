package pgstore

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/models"
)

//SaveSession SaveSession
func (pg *PgStore) SaveSession(session *models.Session) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Create(session)

}

//GetSession GetSession
func (pg *PgStore) GetSession(sessionModel *models.Session) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Preload("User").Where("session = ?", sessionModel.Session).First(sessionModel)

}

//GetSessionByUserID GetSessionByUserID
func (pg *PgStore) GetSessionByUserID(sessionModel *models.Session) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Preload("User").Where("user_id = ?", sessionModel.UserID).First(sessionModel)

}

//UpdateSession UpdateSession
func (pg *PgStore) UpdateSession(session *models.Session) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Model(session).Updates(models.Session{
		Session: session.Session,
	})

}

//DeleteSession DeleteSession
func (pg *PgStore) DeleteSession(session *models.Session) *gorm.DB {

	conn := pg.EstablishConnection()
	defer conn.Close()
	return conn.Unscoped().Delete(models.Session{}, "session = ?", session.Session)

}
