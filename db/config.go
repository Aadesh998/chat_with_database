package db

import (
	"fmt"
	"log"

	"github.com/Aadesh-lab/envloader"
	"github.com/Aadesh-lab/views"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDBConfig() views.DBConfig {
	return views.DBConfig{
		User:     envloader.AppConfig.DBUser,
		Password: envloader.AppConfig.DBPassword,
		Host:     envloader.AppConfig.DBHost,
		Port:     envloader.AppConfig.DBPort,
		Name:     envloader.AppConfig.DBName,
		Schema:   envloader.AppConfig.DBSchema,
	}
}

func GetDSN() string {
	cfg := GetDBConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=require TimeZone=Asia/Kolkata",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.Schema,
	)
	return dsn
}

func InitDB() {
	dsn := GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL DB:", err)
	}
	DB = db
}
