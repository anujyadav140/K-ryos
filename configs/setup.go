package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"koryos/models"
)
var DB *gorm.DB
var err error

func ConnectDB() {

	DB, err = gorm.Open(postgres.Open(envConfig()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	DB.AutoMigrate(&models.Room{})
	DB.AutoMigrate(&models.Message{})

}