package main

import (
	"log"

	config "github.com/edorguez/payment-reminder/configs/account"
	"github.com/edorguez/payment-reminder/internal/account"
	"github.com/edorguez/payment-reminder/internal/account/handlers"
	"github.com/edorguez/payment-reminder/internal/account/models"
	"github.com/edorguez/payment-reminder/internal/account/repository"
	"github.com/edorguez/payment-reminder/internal/account/services"
	"github.com/edorguez/payment-reminder/pkg/database/postgresql"
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

	db, err := postgresql.DBConnection(dbConnection)
	if err != nil {
		return
	}

	models.AutoMigrateModels(db)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes := account.NewRoutes(*userHandler)
	err = routes.Start("0.0.0.0:" + c.Account_Svc_Port)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
