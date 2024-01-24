package database

import (
	"fmt"
	"log"

	"github.com/Pdv2323/login-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var db *gorm.DB
// var Users models.User

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbName     = "myuser"
	dbUser     = "safari"
	dbPassword = "qwerty"
)

func ConnectDB() (*gorm.DB, error) {
	var Users models.User

	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)

	var err error
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&Users)
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
		return nil, err
	}
	log.Println("Connected to the database successfully")
	return db, nil
}
