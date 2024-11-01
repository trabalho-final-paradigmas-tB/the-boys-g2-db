package main

import (
	"backend/database"
	"backend/server"
)

func main() {
	database.ConnectDB()
	server.StartServer()
}
