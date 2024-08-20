package storage

import (
	"core/processing/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DBInit(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.BotConfig{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
