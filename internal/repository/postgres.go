package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToPostgres(host string, user string, password string, dbName string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, 5432)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
