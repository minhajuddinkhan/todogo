package db

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Database Database Interface
type Database interface {
	EstablishConnection() *gorm.DB
	Migrate(models []interface{})
	Initialize(models []interface{})
	SeedDB()
}
