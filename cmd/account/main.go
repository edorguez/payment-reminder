package main

import (
	"github.com/edorguez/payment-reminder/internal/account/handlers"
	"github.com/edorguez/payment-reminder/internal/account/repository"
	"github.com/edorguez/payment-reminder/internal/account/services"
	"github.com/edorguez/payment-reminder/pkg/database/postgresql"
)

func main() {
	db, err := postgresql.DBConnection("")
	if err != nil {
		return
	}

	// Repository Layer
	userRepo := repository.NewUserRepository(db)

	// Service Layer
	userService := services.NewUserService(userRepo)

	// Handler Layer
	userHandler := handlers.NewUserHandler(userService)
}
