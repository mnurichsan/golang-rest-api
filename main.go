package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mnurichsan/golang-rest-api/database"
	"github.com/mnurichsan/golang-rest-api/routes"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()
	routes.SetupRoutes(app)
	port := os.Getenv("PORT")
	app.Listen(port)
}
