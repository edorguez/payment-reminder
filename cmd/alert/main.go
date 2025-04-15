package main

import (
	"log"

	config "github.com/edorguez/payment-reminder/configs/alert"
	"github.com/edorguez/payment-reminder/internal/alert"
	"github.com/edorguez/payment-reminder/internal/alert/handlers"
	"github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/edorguez/payment-reminder/internal/alert/repository"
	"github.com/edorguez/payment-reminder/internal/alert/services"
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

	alertRepo := repository.NewAlertRepository(db)
	alertService := services.NewAlertService(alertRepo)
	alertHandler := handlers.NewAlertHandler(alertService)

	routes := alert.NewRoutes(*alertHandler)
	err = routes.Start("0.0.0.0:" + c.Alert_Svc_Port)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
