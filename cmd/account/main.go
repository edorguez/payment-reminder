package main

import (
	"log"

	"github.com/edorguez/payment-reminder/internal/account"
	"github.com/edorguez/payment-reminder/internal/account/handlers"
	"github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/repository"
	"github.com/edorguez/payment-reminder/internal/account/services"
	"github.com/edorguez/payment-reminder/pkg/database/postgresql"
)

func main() {
	db, err := postgresql.DBConnection("")
	if err != nil {
		return
	}

	models.AutoMigrateModels(db)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes := account.NewRoutes(*userHandler)
	err = routes.Start("some addres")
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
