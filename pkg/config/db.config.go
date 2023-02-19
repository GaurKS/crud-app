package db

import (
	"log"
	"os"

	"github.com/GaurKS/crud-app/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Todo{})
	log.Println("Database migration completed!")
	return db
}