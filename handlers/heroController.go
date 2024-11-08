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

	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	rows, err := database.Db.Query("SELECT CODIGO_HEROI, NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, DATA_NASCIMENTO, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS FROM HEROI")
	if err != nil {
		http.Error(w, "Erro ao listar heróis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var herois []models.Heroi

	for rows.Next() {
		var heroi models.Heroi
		err := rows.Scan(&heroi.CodigoHeroi, &heroi.NomeReal, &heroi.NomeHeroi, &heroi.Sexo, &heroi.AlturaHeroi, &heroi.PesoHeroi, &heroi.DataNascimento, &heroi.LocalNascimento, &heroi.Poderes, &heroi.NivelForca, &heroi.Popularidade, &heroi.Status, &heroi.HistoricoBatalhas)
		if err != nil {
			http.Error(w, "Erro ao escanear herói: "+err.Error(), http.StatusInternalServerError)
			return
		}
		herois = append(herois, heroi)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro na leitura das linhas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(herois)
}

func ListarHeroiPorID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
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

func ModificarHeroi(w http.ResponseWriter, r *http.Request) {
	// Verificar conexão com o banco de dados
	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	// Decodificar o JSON para a estrutura Heroi
	var heroi models.Heroi
	if err := json.NewDecoder(r.Body).Decode(&heroi); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Obter o ID do herói a partir dos parâmetros da URL
	vars := mux.Vars(r)
	heroiIDStr := vars["id"]
	heroiID, err := strconv.Atoi(heroiIDStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Construir a consulta SQL
	query := `UPDATE HEROI
        SET NOME_REAL = $1, NOME_HEROI = $2, SEXO = $3, ALTURA_HEROI = $4, PESO_HEROI = $5, LOCAL_NASCIMENTO = $6, PODERES = $7, NIVEL_FORCA = $8, POPULARIDADE = $9, STATUS = $10, HISTORICO_BATALHAS = $11, DATA_NASCIMENTO = $12
        WHERE CODIGO_HEROI = $13`

	// Executar a consulta com os valores parametrizados
	_, err = database.Db.Exec(query,
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
		heroiID,
	)
	if err != nil {
		http.Error(w, "Erro ao atualizar herói: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder com sucesso
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Herói atualizado com sucesso!"})
}
