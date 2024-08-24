package main

import (
	"log"
	"os"

	"github.com/JongSinister/TeeYai_2024/config"
	// "github.com/JongSinister/TeeYai_2024/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
    // pwd, err := os.Getwd()
    // if err != nil {
    //     panic(err)
    // }

    // Load environment variables from .env file
    if err := godotenv.Load("../config/.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Connect to MongoDB
    config.ConnectDB()
    defer config.DisconnectDB()

    // Init Fiber
    app := fiber.New()

    // Set routes
    // routes.SetupRoutes(app)

    // Start server
    port := os.Getenv("PORT")
    log.Fatal(app.Listen("localhost:" + port))
}
