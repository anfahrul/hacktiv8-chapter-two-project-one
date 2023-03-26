package database

import (
	// "gorm.io/driver/postgres"
	"fmt"
	"hacktiv8-chapter-two-project-one/models/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "pwd"
	dbName   = "hacktiv8_books_2"
	port     = 5433
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(entity.Book{})
}

func GetDB() *gorm.DB {
	return db
}
