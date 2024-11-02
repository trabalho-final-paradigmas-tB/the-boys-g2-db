package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// DB é uma variável global para armazenar a conexão com o banco de dados.
var DB *sql.DB

// Conecta ao SQL Server
func ConnectToSQLServer() {
	connectionString := "sqlserver://grupo_the_boys:adm10@localhost:1433?database=THE_BOYS"

	var err error
	DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão: %s", err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao SQL Server: %s", err.Error())
	}

	fmt.Println("Conexão bem-sucedida com o SQL Server!")
}
