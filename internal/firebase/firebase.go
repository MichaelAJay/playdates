package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var authClient *auth.Client

func InitFirebase() {
	// @TODO - create SA & figure out how to get it into the project correctly - it's done with GOOGLE_APPLICATION_CREDENTIALS in the .env in Node
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Initialize the Auth client
	authClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Auth client: %v", err)
	}
}

func GetAuthClient() *auth.Client {
	return authClient
}
