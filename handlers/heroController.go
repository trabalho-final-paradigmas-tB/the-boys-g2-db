package handlers

import (
	"backend/database"
	"fmt"
	"net/http"
)

var db, _ = database.ConnectDB()

func InserirHeroi(w http.ResponseWriter, r *http.Request) {
	// Exemplo de inserção de dados usando a conexão `database.DB`
	insertQuery := "INSERT INTO herois (NOME_REAL, NOME_HEROI) VALUES (?, ?)"
	res, err := db.Exec(insertQuery, "Clark Kent", "Superman")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Herói inserido com sucesso! Último ID inserido: %d\n", lastInsertID)
}

/*
func AtualizarHeroi(w http.ResponseWriter, r *http.Request) {

	// Exemplo de atualização de dados
	updateQuery := "UPDATE sua_tabela SET coluna1 = ? WHERE id = ?"
	w, err = db.Exec(updateQuery, "novo_valor", lastInsertID)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Número de linhas afetadas: %d\n", rowsAffected)
}

func DeletarHeroi(w http.ResponseWriter, r *http.Request) {

	// Exemplo de exclusão de dados
	deleteQuery := "DELETE FROM sua_tabela WHERE id = ?"
	res, err = db.Exec(deleteQuery, lastInsertID)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Número de linhas excluídas: %d\n", rowsAffected)
}
*/
