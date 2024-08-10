package db

import (
	"backend/db/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ConnStr = "user=postgres password=postgres123 dbname=mydatabase host=postgres port=5432 sslmode=disable"

func SetupDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error Connecting to database %v", err.Error())
		return
	}

	// Automigrate schema
	if err := DB.AutoMigrate(&models.Employee{}, &models.Department{}); err != nil {
		log.Fatalf("Error Migrating Tables: %v", err)
		return
	}

	log.Println("Successfully created and migrated tables")
}
