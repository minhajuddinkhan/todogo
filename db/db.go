package db

import (
	"fmt"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	connStr string
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
	conn, err := gorm.Open(pdb.dialect, pdb.connStr)

	if err != nil {
		panic(err)
	}
	return conn
}

//Migrate Migrate
func (pdb *PostgresDB) Migrate(models []interface{}) {

	conn := pdb.EstablishConnection()
	defer conn.Close()
	for _, model := range models {
		if !conn.HasTable(model) {
			conn.AutoMigrate(model)
		}

	}
	logrus.Info("Database Migrated!")

}

//Initialize Initialize
func (pdb *PostgresDB) Initialize(models []interface{}) {

	pdb.Migrate(models)
	pdb.SeedDB()
	logrus.Info("Database Seeded!")

}

//SeedAll SeedDB
func (pdb *PostgresDB) SeedDB() {

	conn := pdb.EstablishConnection()
	defer conn.Close()
	err := conn.First(&models.User{}, "1").Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			conn.Create(&models.User{
				Name:     "Rameez",
				Address:  "Orangi",
				Password: "123",
			})
		} else {

			fmt.Println(err.Error())

		}

	}

	err = conn.Where("id = 1").First(&models.Todo{}).Error
	if gorm.IsRecordNotFoundError(err) {
		conn.Create(&models.Todo{
			Name:    "Eat Food.",
			Priorty: 1,
			UserID:  1,
		})

	}

}
