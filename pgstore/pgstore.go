package pgstore

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/db"
)

type PgStore struct {
	DB *db.PostgresDB
}

func NewPgStore(db *db.PostgresDB) *PgStore {
	return &PgStore{
		DB: db,
	}

}

//EstablishConnection EstablishConnection
func (p *PgStore) EstablishConnection() *gorm.DB {

	pgConn := p.DB.EstablishConnection()
	//	pgConn.LogMode(true)
	return pgConn

}
