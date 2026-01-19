package envloader

import (
	"log"
	"os"

	"github.com/Aadesh-lab/views"
	"github.com/joho/godotenv"
)

var AppConfig *views.Config

func LoadConfig() {
	env := "developement"
	envFile := ".env." + env

	log.Println(envFile)

	if envFile == "" {
		log.Println("Envloader file is empty")
		os.Exit(1)
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Warning: could not load %s file, falling back to .env: %s", envFile, err)
		err = godotenv.Load()
		if err != nil {
			log.Printf("Warning: could not load .env file either: %s", err)
		}
	}

	AppConfig = &views.Config{
		Version:    os.Getenv("VERSION"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBSchema:   os.Getenv("DB_SCHEMA"),
	}
}
