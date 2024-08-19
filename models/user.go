package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string         `gorm:"unique" json:"username"`
	Password  string         `json:"password"`
	FirstName string         `gorm:"column:first_name" json:"firstname"`
	LastName  string         `gorm:"column:last_name" json:"lastname"`
	Age       int            `json:"age"`
	Roles     string         `json:"roles"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func MigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
