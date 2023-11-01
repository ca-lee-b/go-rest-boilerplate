package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToPostgres() (*gorm.DB, error) {
	dsn := "host=0.0.0.0 user=postgres password=hello dbname=gorm port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
