package config

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitializeFirebase(serviceAccountPath string) {
	opt := option.WithCredentialsFile(serviceAccountPath)

	// Initialize the Firebase app
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	FirebaseApp = app // Assign to the exported variable
	log.Println("Firebase initialized successfully!")
}

func GetFirestoreClient() (*firestore.Client, error) {
	if FirebaseApp == nil {
		return nil, fmt.Errorf("Firebase app not initialized")
	}

	client, err := FirebaseApp.Firestore(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}