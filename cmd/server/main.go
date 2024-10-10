package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/Dating-App-Api/config"
	"github.com/joshua468/Dating-App-Api/internal/routes"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	router := gin.Default()

	// Set up the routes
	routes.SetupRoutes(router)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
