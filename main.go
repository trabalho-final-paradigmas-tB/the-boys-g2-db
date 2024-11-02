package main

import (
	"backend/database"
	"backend/server"
)

func main() {
	// Conecta ao banco de dados
	database.ConnectDB()
	defer database.DB.Close() // Fecha a conexão ao final

	server.StartServer()
}
