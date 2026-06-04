package main

import (
	"log"

	"github.com/joho/godotenv"

	"ticket-booking/internal/database"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Migration completed")
}