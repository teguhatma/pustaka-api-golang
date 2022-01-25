package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConn() *gorm.DB {
	db_host := EnvVariable("DB_HOST")
	db_user := EnvVariable("DB_USER")
	db_password := EnvVariable("DB_PASSWORD")
	db_name := EnvVariable("DB_NAME")
	db_port := EnvVariable("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error")
	}

	fmt.Println("Database connection success")
	return db
}
