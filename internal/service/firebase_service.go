package service

import (
	"context"
	"log"

	"firebase.google.com/go/v4/messaging"
	"toba-lapor-backend/internal/config"
)

type FirebaseService interface {
	SendPushNotification(token string, title string, body string) error
}

type firebaseService struct{}

func NewFirebaseService() FirebaseService {
	return &firebaseService{}
}

func (s *firebaseService) SendPushNotification(token string, title string, body string) error {
	if config.FirebaseApp == nil {
		log.Println("Push Notification tertahan: Firebase App belum terinisialisasi.")
		return nil
	}

	if token == "" {
		return nil // Jangan kirim kalau tidak ada token
	}

	client, err := config.FirebaseApp.Messaging(context.Background())
	if err != nil {
		log.Printf("Gagal mendapatkan Messaging client: %v\n", err)
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	}

	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Printf("Gagal mengirim notifikasi ke %s: %v\n", token, err)
		return err
	}

	log.Println("Sukses mengirim Push Notification:", response)
	return nil
}
