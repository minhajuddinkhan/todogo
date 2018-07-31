package models

import (
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Session string
	User    User `gorm:"foreignkey:UserID"`
	UserID  uint
}
