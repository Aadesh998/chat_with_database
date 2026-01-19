package main

import (
	"log"

	"github.com/Aadesh-lab/db"
	dbfunction "github.com/Aadesh-lab/db_function"
	"github.com/Aadesh-lab/envloader"
)

func main() {
	envloader.LoadConfig()
	db.InitDB()

	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}

	schema, err := dbfunction.GetFullSchema(sqlDB, envloader.AppConfig.DBSchema)
	if err != nil {
		log.Fatal(err)
	}

	dbfunction.PrintSchema(schema)
}
