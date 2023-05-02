package database

import (
	"fmt"
	"log"
	model "middleware/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "bondan2000"
	dbName   = "ecommerce"
	port     = 5432
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

	db.Debug().AutoMigrate(model.User{}, model.Product{})
}

func GetDB() *gorm.DB {
	return db
}
