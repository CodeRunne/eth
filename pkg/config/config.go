package config

import (
	"github.com/joho/godotenv"
	"github.com/infinityethback/pkg/database"
	"github.com/infinityethback/pkg/models"
	"log"
	"os"
)

func init() {
	loadEnv()

	dbData := map[string]string {
		"host": os.Getenv("DB_HOST"),
        "port": os.Getenv("DB_PORT"),
        "user": os.Getenv("DB_USER"),
        "pass": os.Getenv("DB_PASS"),
        "name": os.Getenv("DB_NAME"),
	}

	database.Connect(dbData)
	database.GetDB().AutoMigrate(&models.User{}, &models.Referral{}, &models.Quote{}, &models.Social{})
}

func loadEnv() {
	err := godotenv.Load("../.env")
    if err!= nil {
        log.Fatal(err)
    }
}