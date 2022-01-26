package utils

import (
	"pustaka-api/repository"

	"fmt"

	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(&repository.Book{})
	fmt.Println("Table Book Created.")
	db.AutoMigrate(&repository.User{})
	fmt.Println("Table User Created.")
}
