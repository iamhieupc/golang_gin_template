package database


import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB) {
	dsn := "host=localhost user=postgres password=postgres dbname=golang_gin_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
