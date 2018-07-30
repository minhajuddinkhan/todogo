package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100)`
	Priorty int
	User    User `gorm:"foreignkey:UserID"`
	UserID  int
}
