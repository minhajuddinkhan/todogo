package db

import (
	"fmt"

	gorm "github.com/jinzhu/gorm"
	//SQLITE dialect support
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/sirupsen/logrus"
)

//SqliteDB SqliteDB instance
type SqliteDB struct {
	ConnectionStr string
	dialect       string
}

//NewSqliteDB Creates a new SqliteDB Instance
func NewSqliteDB(conf *config.Configuration) Database {
	return &SqliteDB{
		ConnectionStr: fmt.Sprintf("%s", conf.Db.VolumePath),
		dialect:       "sqlite3",
	}
}

//EstablishConnection EstablishConnection
func (pdb *SqliteDB) EstablishConnection() *gorm.DB {
	conn, err := gorm.Open(pdb.dialect, pdb.ConnectionStr)
	if err != nil {
		panic(err)
	}
	return conn
}

//Migrate Migrate
func (pdb *SqliteDB) Migrate(models []interface{}) {

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
func (pdb *SqliteDB) Initialize(models []interface{}) {

	pdb.Migrate(models)
	pdb.SeedDB()
	logrus.Info("Database Seeded!")

}

//SeedDB SeedDB
func (pdb *SqliteDB) SeedDB() {

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
