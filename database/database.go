package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func connect() error {
	dsn := os.Getenv("DATABASE_URL")
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db = _db

	log.Println("Database connection established")
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		if err := connect(); err != nil {
			panic(err)
		}
	}

	return db
}
