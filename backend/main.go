package main

import (
	routes "franchise-web/config"
	"franchise-web/config/initializers"
	"log"
	"os"
)

func init() {
	log.Println("== Let's Go!! ==")
	appIP := os.Getenv("HOST")
	appPort := os.Getenv("PORT")
	log.Printf("appIP: %s, appPort: %s", appIP, appPort)
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
