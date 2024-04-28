package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"

	config2 "backend/config"
)

var App *firebase.App
var FirebaseAuth *auth.Client

// InitFirebase initializes the Firebase app and Firebase Auth client
// It should be called before any other Firebase function.
func InitFirebase() error {
	var err error

	config, err := config2.LoadConfig()
	if err != nil {
		return err
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(config.FirebaseConfigDir)
	App, err = firebase.NewApp(ctx, nil, opt)
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
