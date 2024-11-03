package main

import (
	"backend/database"
	"backend/server"
	"log"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}
	defer database.Db.Close()
	server.StartServer()
}
