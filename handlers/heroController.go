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
	(CODIGO_HEROI, NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS, DATA_NASCIMENTO)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING CODIGO_HEROI`

	var lastInsertID int
	err := database.Db.QueryRow(insertQuery,
		heroi.ID,
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
	query := "SELECT NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, DATA_NASCIMENTO, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS FROM herois"

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
		http.Error(w, "ID invalido "+err.Error(), http.StatusBadRequest)
		return
	}

	deleteQuery := "DELETE FROM herois WHERE CODIGO_HEROI = $1"
	res, err := database.Db.Exec(deleteQuery, heroid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao tentar deletar herói : %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao verificar linha afetadas: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Número de linhas excluídas: %d\n", rowsAffected)

	if rowsAffected == 0 {
		http.Error(w, "Herói  não encontrado ", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"mensagem": "herois foi deletado com sucesso",
		"status":   "sucesso ",
		"code":     200,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Erro ao mandar mensagem ", http.StatusInternalServerError)

	}

}
