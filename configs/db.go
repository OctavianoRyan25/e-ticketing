package configs

import (
	"github.com/OctavianoRyan25/e-ticketing/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=1234 dbname=eticket port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
}

func AutoMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Terminal{})
}
