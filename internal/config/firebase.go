package config

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func ConnectFirebase() {
	// Letakkan file JSON Anda di root folder project
	serviceAccountKeyPath := "firebase-service-account.json"

	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Warning: Gagal inisialisasi Firebase (file JSON %s mungkin belum ada): %v", serviceAccountKeyPath, err)
		return
	}

	FirebaseApp = app
	fmt.Println("Firebase Admin SDK berhasil diinisialisasi")
}
