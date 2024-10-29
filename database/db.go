package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "aqui fica o link / chave do mysql - nao sei se vai ser online ou localhost msm"
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
