package main

import (
	"effective-mobile/online-subscriptions/internal/database"
	"effective-mobile/online-subscriptions/internal/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := database.Connect()

	err := db.AutoMigrate(&models.Subscription{})
	if err != nil {
		panic(err)
	}

	log.Println("database migrated")
}
