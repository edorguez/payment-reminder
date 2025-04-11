package main

import (
	"log"

	"github.com/edorguez/payment-reminder/internal/alert"
	"github.com/edorguez/payment-reminder/internal/alert/handlers"
	"github.com/edorguez/payment-reminder/internal/alert/models"
	"github.com/edorguez/payment-reminder/internal/alert/repository"
	"github.com/edorguez/payment-reminder/internal/alert/services"
	"github.com/edorguez/payment-reminder/pkg/database/postgresql"
)

func main() {
	db, err := postgresql.DBConnection("")
	if err != nil {
		return
	}

	models.AutoMigrateModels(db)

	alertRepo := repository.NewAlertRepository(db)
	alertService := services.NewAlertService(alertRepo)
	alertHandler := handlers.NewAlertHandler(alertService)

	routes := alert.NewRoutes(*alertHandler)
	err = routes.Start("some addres")
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
