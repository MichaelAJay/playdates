package main

import (
	"playdates/internal/firebase"
	"playdates/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialization
	firebase.InitFirebase()

	// Create a new Gin router
	router := gin.New()

	// Initialize routes
	routes.InitRoutes(router)

	// Start the server
	router.Run(":8080")
}
