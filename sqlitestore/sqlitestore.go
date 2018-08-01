package sqlitestore

import (
	"github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/db"
	"github.com/minhajuddinkhan/todogo/store"
)

//SqliteStore Sqlite store instance
type SqliteStore struct {
	DB db.Database
}

//NewSqliteStore Creates a new Sqlite store instance.
func NewSqliteStore(database db.Database) store.Store {

	return &SqliteStore{
		DB: database,
	}
}

//EstablishConnection EstablishConnection
func (p *SqliteStore) EstablishConnection() *gorm.DB {

	pgConn := p.DB.EstablishConnection()
	//	pgConn.LogMode(true)
	return pgConn

}
