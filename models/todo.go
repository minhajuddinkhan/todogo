package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100)" json:"name,omitempty"`
	Priorty int    `json:"priority,omitempty"`
	User    User   `gorm:"foreignkey:UserID"`
	UserID  uint   `json:"userId,omitempty"`
}
