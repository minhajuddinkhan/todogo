package seeds

import (
	"github.com/jinzhu/gorm"
)

//SeedAll SeedAll
func SeedAll(db *gorm.DB, modelInstances []interface{}) {

	for _, modelInstance := range modelInstances {
		db.Create(modelInstance)
	}
}

func seedOne(db *gorm.DB, modelInstance interface{}) {
	db.Create(modelInstance)
}
