package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "host=10.60.65.122 user=todolist password=todolist dbname=todolist port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
