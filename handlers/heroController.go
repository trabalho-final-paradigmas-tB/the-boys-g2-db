package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InserirHeroi(w http.ResponseWriter, r *http.Request) {

	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	var heroi models.Heroi
	if err := json.NewDecoder(r.Body).Decode(&heroi); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	insertQuery := `INSERT INTO HEROI
	(NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS, DATA_NASCIMENTO)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING CODIGO_HEROI`

	var lastInsertID int
	err := database.Db.QueryRow(insertQuery,
		heroi.NomeReal,
		heroi.NomeHeroi,
		heroi.Sexo,
		heroi.AlturaHeroi,
		heroi.PesoHeroi,
		heroi.LocalNascimento,
		heroi.Poderes,
		heroi.NivelForca,
		heroi.Popularidade,
		heroi.Status,
		heroi.HistoricoBatalhas,
		heroi.DataNascimento,
	).Scan(&lastInsertID)

	if err != nil {
		http.Error(w, "Erro ao inserir herói: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Herói inserido com sucesso! Último ID inserido: %d\n", lastInsertID)
}

func ListarHerois(w http.ResponseWriter, r *http.Request) {
	query := "SELECT NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, DATA_NASCIMENTO, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS FROM HEROI"

	rows, err := database.Db.Query(query)
	if err != nil {
		http.Error(w, "Erro ao consultar heróis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var herois []models.Heroi
	for rows.Next() {
		var heroi models.Heroi
		if err := rows.Scan(&heroi.NomeReal, &heroi.NomeHeroi, &heroi.Sexo, &heroi.AlturaHeroi, &heroi.PesoHeroi, &heroi.DataNascimento, &heroi.LocalNascimento, &heroi.Poderes, &heroi.NivelForca, &heroi.Popularidade, &heroi.Status, &heroi.HistoricoBatalhas); err != nil {
			http.Error(w, "Erro ao escanear herói: "+err.Error(), http.StatusInternalServerError)
			return
		}
		herois = append(herois, heroi)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(herois)
}

/*func AtualizarHeroi(w http.ResponseWriter, r *http.Request) {

	// Exemplo de atualização de dados
	updateQuery := "UPDATE sua_tabela SET coluna1 = ? WHERE id = ?"
	w, err := db.Exec(updateQuery, "novo_valor", lastInsertID)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Número de linhas afetadas: %d\n", rowsAffected)
}*/

func ListarHeroiPorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	query := `SELECT NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, DATA_NASCIMENTO,
	LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS
	FROM HEROI WHERE CODIGO_HEROI = $1`
	var heroi models.Heroi
	err = database.Db.QueryRow(query, id).Scan(
		&heroi.NomeReal,
		&heroi.NomeHeroi,
		&heroi.Sexo,
		&heroi.AlturaHeroi,
		&heroi.PesoHeroi,
		&heroi.DataNascimento,
		&heroi.LocalNascimento,
		&heroi.Poderes,
		&heroi.NivelForca,
		&heroi.Popularidade,
		&heroi.Status,
		&heroi.HistoricoBatalhas,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Herói não encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao consultar herói: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroi)
}

func DeletarHeroi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	heroiIDStr := vars["id"]

	heroid, err := strconv.Atoi(heroiIDStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// consultar o herói antes de deletá-lo para obter o nome

	var nomeHeroi string
	query := "SELECT NOME_HEROI FROM HEROI WHERE CODIGO_HEROI = $1"
	err = database.Db.QueryRow(query, heroid).Scan(&nomeHeroi)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Herói não encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao buscar herói: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// executar o comando DELETE

	deleteQuery := "DELETE FROM HEROI WHERE CODIGO_HEROI = $1"
	res, err := database.Db.Exec(deleteQuery, heroid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao tentar deletar herói: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Erro ao deletar o herói ou herói não encontrado", http.StatusNotFound)
		return
	}

	// retornar o ID e o nome do herói que foi deletado em vez só de 'herois'

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"mensagem":     "Herói deletado com sucesso",
		"status":       "sucesso",
		"codigo_heroi": heroid,
		"nome_heroi":   nomeHeroi,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao enviar a resposta", http.StatusInternalServerError)
	}
}
