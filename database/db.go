package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() error {

	dsn := "root:HaXQvofgJIsdoWRVTjhmXgNjFmggirNj@tcp(junction.proxy.rlwy.net:38833)/railway" // railway temporariamente

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
