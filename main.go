package main

import (
	"backend/database"
	"log"
)

func main() {

	database, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}
	defer database.Close()

	log.Println("Aplicação conectada e pronta para uso!")
}
