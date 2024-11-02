package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	// Ajuste do DSN para o formato aceito pelo driver
	dsn := "root:vJzVmyIdRQCRKGfApRbhySYziJxYLZCa@tcp(autorack.proxy.rlwy.net:52440)/railway"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão com o banco de dados: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao verificar a conexão com o banco de dados: %w", err)
	}
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return db, nil
}
