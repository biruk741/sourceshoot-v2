package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var App *firebase.App
var FirebaseAuth *auth.Client

// InitFirebase initializes the Firebase app and Firebase Auth client
// It should be called before any other Firebase function.
func InitFirebase() error {
	var err error

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &firebase.Config{
		ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
	}

	creds := os.Getenv("FIREBASE_CREDENTIALS")
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(creds))
	App, err = firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v\n", err)
	}

	// Initialize Firebase Auth client
	FirebaseAuth, err = App.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v\n", err)
		return err
	}

	return nil
}
