package database

import (
	"fmt"
	"log"
	"pustaka-api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConn() *gorm.DB {
	db_host := utils.EnvVariable("DB_HOST")
	db_user := utils.EnvVariable("DB_USER")
	db_password := utils.EnvVariable("DB_PASSWORD")
	db_name := utils.EnvVariable("DB_NAME")
	db_port := utils.EnvVariable("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error")
	}

	return db
}
