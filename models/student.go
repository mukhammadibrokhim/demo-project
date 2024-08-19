package models

import (
	"demo-project/config"
	"fmt"
)

type Student struct {
	ID    uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Email string `json:"email"`
}

func MigrateStudent() {
	err := config.DB.AutoMigrate(&Student{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
