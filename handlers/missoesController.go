package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
)

func inserirMissao(w http.ResponseWriter, r *http.Request) {

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

	var rank_escolhida string

	switch missoes.Rank_Escolhida {
	case "SS":
		rank_escolhida = missoes.Rank_SS
	case "S":
		rank_escolhida = missoes.Rank_S
	case "A":
		rank_escolhida = missoes.Rank_A
	case "B":
		rank_escolhida = missoes.Rank_B
	case "C":
		rank_escolhida = missoes.Rank_C
	case "D":
		rank_escolhida = missoes.Rank_D
	case "E":
		rank_escolhida = missoes.Rank_E
	default:
		http.Error(w, "Rank inválido", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO missoes (nome, descricao, classificacao, rank_escolhido)
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	var id int

	err = database.Db.QueryRow(
		query,
		missoes.Nome,
		missoes.Descrição,
		missoes.Classificação,
		rank_escolhida,
	).Scan(&id)

	if err != nil {
		http.Error(w, "Erro ao inserir missão", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"mensagem":       "Missão inserida com sucesso",
		"status":         "sucesso",
		"missao":         missoes.Nome,
		"descrição":      missoes.Descrição,
		"rank_escolhido": rank_escolhida,
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

	rows, err := database.Db.Query("SELECT nome, descricao, classificacao, rank_escolhido FROM missoes")
	if err != nil {
		http.Error(w, "Erro ao listar missões: ", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var missoes []models.Missoes

	for rows.Next() {
		var missao models.Missoes
		err := rows.Scan(&missao.Nome, &missao.Descrição, &missao.Classificação, &missao.Rank_Escolhida)
		if err != nil {
			http.Error(w, "Erro ao escanear missão: ", http.StatusInternalServerError)
			return
		}
		missoes = append(missoes, missao)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro na leitura das linhas: ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(missoes)
}
