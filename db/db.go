package db

import (
	"fmt"

	"github.com/Pdv2323/Login-Auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbName     = "postgres"
	dbUser     = "postgres"
	dbPassword = "123"
)

func ConnectDB() (db *gorm.DB, err error) {
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)
	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(&models.User{})
	return
}
