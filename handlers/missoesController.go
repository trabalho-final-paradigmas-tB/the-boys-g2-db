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

func InserirMissao(w http.ResponseWriter, r *http.Request) {

	if database.Db == nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}

	var missoes models.Missoes
	err := json.NewDecoder(r.Body).Decode(&missoes)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
		return
	}

	if missoes.Dificuldade < 1 || missoes.Dificuldade > 10 {
		http.Error(w, "Dificuldade invalida", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO missoes (nome, descricao, classificacao, dificuldade)
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	var id int

	err = database.Db.QueryRow(
		query,
		missoes.Nome,
		missoes.Descrição,
		missoes.Classificação,
		missoes.Dificuldade,
	).Scan(&id)

	if err != nil {
		http.Error(w, "Erro ao inserir missão", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"mensagem":              "Missão inserida com sucesso",
		"status":                "sucesso",
		"missao":                missoes.Nome,
		"descrição":             missoes.Descrição,
		"dificuldade_escolhida": missoes.Dificuldade,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao enviar a resposta", http.StatusInternalServerError)
	}
}

func ListadeMissões(w http.ResponseWriter, r *http.Request) {

	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	rows, err := database.Db.Query("SELECT nome, descricao, classificacao, dificuldade FROM missoes")
	if err != nil {
		http.Error(w, "Erro ao listar missões", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var missoes []models.Missoes

	for rows.Next() {
		var missao models.Missoes
		err := rows.Scan(&missao.Nome, &missao.Descrição, &missao.Classificação, &missao.Dificuldade)
		if err != nil {
			http.Error(w, "Erro ao escanear missão", http.StatusInternalServerError)
			return
		}
		missoes = append(missoes, missao)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro na leitura das linhas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(missoes)
}

func DeletarMissão(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	missaoIDStr := vars["id"]

	missaoID, err := strconv.Atoi(missaoIDStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	var nomeMissao string
	query := "SELECT nome FROM missoes WHERE id = $1"
	err = database.Db.QueryRow(query, missaoID).Scan(&nomeMissao)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Missão não encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao buscar missão: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	deleteQuery := "DELETE FROM missoes WHERE id = $1"
	res, err := database.Db.Exec(deleteQuery, missaoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao tentar deletar missão: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Erro ao deletar a missão ou missão não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"mensagem":      "Missão deletada com sucesso",
		"status":        "sucesso",
		"codigo_missao": missaoID,
		"nome_missao":   nomeMissao,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao enviar a resposta", http.StatusInternalServerError)
	}
}

func ModificarMissao(w http.ResponseWriter, r *http.Request) {

	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	var missao models.Missoes
	if err := json.NewDecoder(r.Body).Decode(&missao); err != nil {
		http.Error(w, "Erro ao decodificar JSON: ", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	missaoIDStr := vars["id"]
	missaoID, err := strconv.Atoi(missaoIDStr)
	if err != nil {
		http.Error(w, "ID inválido: ", http.StatusBadRequest)
		return
	}

	query := `UPDATE missoes
        SET nome = $1, descricao = $2, classificacao = $3, dificuldade = $4
        WHERE id = $5`

	if missao.Dificuldade < 1 || missao.Dificuldade > 10 {
		http.Error(w, "Dificuldade Invalida", http.StatusBadRequest)
		return

	}

	_, err = database.Db.Exec(query,
		missao.Nome,
		missao.Descrição,
		missao.Classificação,
		missao.Dificuldade,
		missaoID,
	)
	if err != nil {
		http.Error(w, "Erro ao atualizar missão: ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensagem": "Missão atualizada com sucesso!"})
}
