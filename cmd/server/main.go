package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Aadesh-lab/db"
	dbfunction "github.com/Aadesh-lab/db_function"
	"github.com/Aadesh-lab/envloader"
	"github.com/Aadesh-lab/services"
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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your query: ")
		userQuery, _ := reader.ReadString('\n')
		userQuery = strings.TrimSpace(userQuery)

		if userQuery == "exit" {
			break
		}

		schemaString := dbfunction.SchemaToString(schema)
		sqlQuery, err := services.LLMCall(schemaString, userQuery)
		if err != nil {
			log.Printf("Error getting SQL from LLM: %v", err)
			continue
		}
		if services.ExcecuteSQL(sqlDB, sqlQuery) {
			fmt.Println("Query executed successfully")

		} else {
			fmt.Println("Failed to execute query")
		}

	}

}
