package postgresql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(DSN string) (*gorm.DB, error) {
	DB, error := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil || DB == nil {
		log.Fatal(error)
	} else {
		log.Println("- Postgres DB Connected")
	}
	return DB, error
}
