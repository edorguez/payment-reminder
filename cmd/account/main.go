package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	config "github.com/edorguez/payment-reminder/configs/account"
	"github.com/edorguez/payment-reminder/internal/account"
	"github.com/edorguez/payment-reminder/internal/account/handlers"
	"github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/repository"
	"github.com/edorguez/payment-reminder/internal/account/services"
	"github.com/edorguez/payment-reminder/pkg/database/postgresql"
	"github.com/edorguez/payment-reminder/pkg/kafka"
	"google.golang.org/api/option"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error reading config file: ", err)
		return
	}

	var dbConnection string
	if c.Environment == "production" {
		dbConnection = c.DB_Source_Production
	} else {
		dbConnection = c.DB_Source_Development
	}

	// Start Portgresql
	db, err := postgresql.DBConnection(dbConnection)
	if err != nil {
		return
	}

	// Use the path to your service account JSON
	opt := option.WithCredentialsFile("../../firebase-adminsdk.json")

	// Initialize Firebae the app
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firbase app: %v\n", err)
		return
	}

	// Get Firebase Auth client
	firebaseClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Firebase Auth client: %v\n", err)
		return
	}

	// Start Kafka producer
	userProducer, err := kafka.NewProducer([]string{"localhost:29092"}, "users")
	if err != nil {
		panic(err)
	}
	defer userProducer.Close()

	// Migrate GORM models
	models.AutoMigrateModels(db)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo, firebaseClient, userProducer)
	userHandler := handlers.NewUserHandler(userService)

	// Start account routes
	routes := account.NewRoutes(*userHandler)
	err = routes.Start("0.0.0.0:" + c.Account_Svc_Port)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
