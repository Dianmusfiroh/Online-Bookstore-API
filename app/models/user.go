package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
    Name      string    `gorm:"type:varchar(100)" json:"name"`
    Email     string    `gorm:"type:varchar(100);unique" json:"email"`
    Password  string    `gorm:"type:varchar(255)" json:"-"`
    Role      string    `gorm:"type:enum('user','admin');default:'user'" json:"role"`
}