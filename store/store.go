package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PgStore struct {
	Dialect          string
	ConnectionString string
	Connetion        *gorm.DB
}

func NewPgStore(connStr string) *PgStore {
	return &PgStore{
		Dialect:          "postgres",
		ConnectionString: connStr,
	}

}

//EstablishConnection EstablishConnection
func (p *PgStore) EstablishConnection() *gorm.DB {

	fmt.Println("p.ConnectionString", p.ConnectionString)
	pgConn, err := gorm.Open(p.Dialect, p.ConnectionString)
	if err != nil {
		panic(err)
	}
	return pgConn

}
