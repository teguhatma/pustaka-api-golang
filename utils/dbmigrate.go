package utils

import "pustaka-api/repository"

func DatabaseMigration() {
	db := DatabaseConn()

	db.AutoMigrate(&repository.Book{})
}
