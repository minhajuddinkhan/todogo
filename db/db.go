package db

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/minhajuddinkhan/todogo/models"
)

type PostgresDB struct {
	connStr string
	Conn    *gorm.DB
	dialect string
}

//NewPostgresDB NewPostgresDB
func NewPostgresDB(conn string, dialect string) *PostgresDB {
	return &PostgresDB{
		connStr: conn,
		dialect: dialect,
	}
}

//EstablishConnection EstablishConnection
func (pdb *PostgresDB) EstablishConnection() *gorm.DB {
	var err error
	pdb.Conn, err = gorm.Open(pdb.dialect, pdb.connStr)
	defer pdb.Conn.Close()
	if err != nil {
		panic(err)
	}
	return pdb.Conn
}

//Migrate Migrate
func (pdb *PostgresDB) Migrate(models []interface{}) {

	for _, model := range models {
		if !pdb.Conn.HasTable(model) {
			pdb.Conn.AutoMigrate(model)
		}

	}
}

//SeedAll SeedDB
func (pdb *PostgresDB) SeedDB() {

	pdb.Conn.Create(&models.User{
		Name:     "Rameez",
		Address:  "Orangi",
		Password: "123",
	})
	pdb.Conn.Create(&models.Todo{
		Name:    "Eat Food.",
		Priorty: 1,
		UserID:  1,
	})

}
