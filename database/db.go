package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {

	/*
		Para testar o banco de dados por enquanto, criem um bd localmente msm e mude o seguinte link: ' root:senha_do_root@tcp(localhost:3306)/meu_banco '
	*/

	var err error
	dsn := "root:ceub123456@tcp(localhost:3306)/teste_sql"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	return db, nil
}
