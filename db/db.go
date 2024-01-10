package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbName     = "mydb"
	dbUser     = "myuser"
	dbPassword = "mypassword"
)

func ConnectDB() (db *gorm.DB, err error) {
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)
	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
