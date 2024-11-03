package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() error {
	user := "root"
	password := "HaXQvofgJIsdoWRVTjhmXgNjFmggirNj"
	host := "junction.proxy.rlwy.net"
	port := "38833"
	dbName := "railway"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao abrir a conexão com o banco de dados: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("erro ao verificar a conexão com o banco de dados: %w", err)
	}
	Db = db
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return nil
}
