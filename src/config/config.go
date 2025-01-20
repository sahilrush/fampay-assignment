package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sahilrush/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construct the DSN (Data Source Name) string from environment variables
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE") +
		" TimeZone=" + os.Getenv("DB_TIMEZONE")

	// Open a connection to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connected successfully!")

	// Migrate the schema (create or update the `videos` table)
	if err := DB.AutoMigrate(&models.Video{}); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Database migrated successfully!")
}
