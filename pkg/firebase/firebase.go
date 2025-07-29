package firebase

import (
	"context"
	"log"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var (
	firebaseClient *auth.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		opt := option.WithCredentialsFile("../../firebase-adminsdk.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("firebase.NewApp: %v", err)
		}
		firebaseClient, err = app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Firebase Auth client: %v\n", err)
			return
		}
	})
}

func Client() *auth.Client { return firebaseClient }
