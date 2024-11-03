package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() error {
	dsn := "postgres://postgres:ZFeBjvz0qvwvGb2H@anciently-native-tody.data-1.use1.tembo.io:5432/postgres"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("erro ao abrir a conexão com o banco de dados: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("erro ao verificar a conexão com o banco de dados: %w", err)
	}
	Db = db
	log.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")
	return nil
}
