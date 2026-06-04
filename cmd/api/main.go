package main

import (
	"log"

	"ticket-booking/internal/database"
	"ticket-booking/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed load env")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}