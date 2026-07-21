package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"toko-online-backend/config"
	"toko-online-backend/routes"
)

func main() {
	log.Println("Starting server...")
	
	// Connect to database
	config.ConnectDatabase()

	// Create Gin router
	r := gin.Default()

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Ping endpoint for testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Setup product routes
	routes.SetupProductRoutes(r)

	// Setup order routes
	routes.SetupOrderRoutes(r)

	// Start server on port 8080
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
