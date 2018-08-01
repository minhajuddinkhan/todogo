package db

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	gorm "github.com/jinzhu/gorm"
	"github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/models"

	"github.com/sirupsen/logrus"
)

type PostgresDB struct {
	ConnectionStr string
	dialect       string
}

//NewPostgresDB Creates a new postgres database
func NewPostgresDB(conf *config.Configuration) Database {
	return &PostgresDB{
		ConnectionStr: fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"+"",
			conf.Db.Host, conf.Db.Port, conf.Db.Username, conf.Db.Name, conf.Db.Password),
		dialect: "postgres",
	}
}

//EstablishConnection EstablishConnection
func (pdb *PostgresDB) EstablishConnection() *gorm.DB {
	conn, err := gorm.Open(pdb.dialect, pdb.ConnectionStr)
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

//SeedDB SeedDB
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
