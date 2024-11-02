package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func InserirHeroi(w http.ResponseWriter, r *http.Request) {
	var heroi models.Heroi
	if err := json.NewDecoder(r.Body).Decode(&heroi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertQuery := `INSERT INTO herois (NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, DATA_NASCIMENTO, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := database.DB.Exec(insertQuery,
		heroi.NomeReal,
		heroi.NomeHeroi,
		heroi.Sexo,
		heroi.AlturaHeroi,
		heroi.PesoHeroi,
		heroi.DataNascimento,
		heroi.LocalNascimento,
		heroi.Poderes,
		heroi.NivelForca,
		heroi.Popularidade,
		heroi.Status,
		heroi.HistoricoBatalhas)

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
