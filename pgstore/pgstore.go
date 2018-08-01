package pgstore

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/db"
	"github.com/minhajuddinkhan/todogo/store"
)

type PgStore struct {
	DB db.Database
}

func NewPgStore(database db.Database) store.Store {
	return &PgStore{
		DB: database,
	}

}

//EstablishConnection EstablishConnection
func (p *PgStore) EstablishConnection() *gorm.DB {

	pgConn := p.DB.EstablishConnection()
	//	pgConn.LogMode(true)
	return pgConn

}
