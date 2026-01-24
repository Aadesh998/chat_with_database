package main

import (
	"log"
	"net/http"

	"github.com/Aadesh-lab/db"
	dbfunction "github.com/Aadesh-lab/db_function"
	"github.com/Aadesh-lab/envloader"
	"github.com/Aadesh-lab/services"
	"github.com/Aadesh-lab/views"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request, schemaStr string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}
	defer conn.Close()

	sqlDB, _ := db.DB.DB()

	for {
		var msg map[string]string
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		userQuery := msg["query"]

		sqlQuery, err := services.LLMCall(schemaStr, userQuery)
		if err != nil {
			conn.WriteJSON(views.WSResponse{Type: "ERROR", Payload: err.Error()})
			continue
		}

		conn.WriteJSON(views.WSResponse{Type: "SQL", Payload: sqlQuery})

		data, err := services.ExecuteForUI(sqlDB, sqlQuery)
		if err != nil {
			conn.WriteJSON(views.WSResponse{Type: "ERROR", Payload: err.Error()})
			continue
		}

		conn.WriteJSON(views.WSResponse{Type: "DATA", Payload: data})
	}
}

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

	schemaStr := dbfunction.SchemaToString(schema)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(w, r, schemaStr)
	})

	log.Println("Server started on :http://localhost:3900/")
	log.Fatal(http.ListenAndServe(":3900", nil))
}
