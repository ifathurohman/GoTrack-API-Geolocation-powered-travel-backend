package databases

import (
	"fmt"
	"os"
	"travel_backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	Database *gorm.DB
}

func init() {
	// Load environment variables from .env file if not in production
	if os.Getenv("ENV") != "Production" {
		if err := godotenv.Load(".env"); err != nil {
			fmt.Println("Error loading .env file")
		}
	}
}

func SetDatabase() (*PostgreSQL, error) {
	// Fetch environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct connection string
	dbConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Make migrations if necessary
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Location{})
	db.AutoMigrate(&models.Image{})
	db.AutoMigrate(&models.Pin{})

	return &PostgreSQL{
		Database: db,
	}, nil
}
