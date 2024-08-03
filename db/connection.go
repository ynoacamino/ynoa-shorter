package db

import (
	"log"

	"github.com/ynoacamino/ynoa-shorter/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN string = "host=localhost user=ynoacamino password=11yenaro11 dbname=gorm port=5432"

var DB *gorm.DB

func DBconnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
}

func DBMigrate() {
	DB.AutoMigrate(&models.Link{})
}
