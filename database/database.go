package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	port = 5432
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to database")
			return
		}

		log.Fatal(fmt.Errorf("attemp %s: error connecting to database: %w", i, err))
		time.Sleep(1 * time.Second)
	}

	log.Fatal("Unable to connect to the database")
}
