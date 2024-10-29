package main

import (
	routes "franchise-web/config"
	"franchise-web/config/initializers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("== Let's Go!! ==")
	appIP := os.Getenv("HOST")
	appPort := os.Getenv("PORT")
	log.Printf("appIP: %s, appPort: %s", appIP, appPort)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	initializers.ConnectToDB()
}

// @title Web API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
	routes := routes.InitRouter()
	if err := routes.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
