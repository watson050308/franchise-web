package main

import (
	routes "franchise-web/config"
	"log"
)

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
