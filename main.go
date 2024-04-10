package main

import (
	"context"
	"log"
	"os"
	"playdates/internal/firebase"
	"playdates/internal/routes"
	"playdates/internal/secretmanager"
	"playdates/internal/security"

	"github.com/gin-gonic/gin"
)

func main() {
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	if projectID == "" {
		log.Fatal("GOOGLE_PROJECT_ID environment variable is not set")
	}

	// Initialize Secret Manager
	secretManagerClient, err := secretmanager.NewSecretManager(context.Background(), projectID)
	if err != nil {
		log.Fatal(err)
	}

	// How will I access the security instance
	// I think there's some mechanism by which I can ensure that the security instance is only initialized once, but I can use this method to get the singleton every time
	_, err = security.NewSecurity(secretManagerClient, projectID)
	if err != nil {
		log.Fatal(err)
	}

	// Initialization
	firebase.InitFirebase()

	// Create a new Gin router
	router := gin.New()

	// Initialize routes
	routes.InitRoutes(router)

	// Start the server
	router.Run(":8080")
}
