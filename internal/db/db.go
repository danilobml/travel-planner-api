package db

import (
	"log"
	"os"

	"github.com/danilobml/travel-planner-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		log.Fatalln("No POSTGRES_URL env variable provided")
	}

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to initialize DB", err)
	}

	db.AutoMigrate(&models.Plan{})

	return db
}
