package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
)

func inserirMissao(w http.ResponseWriter, r *http.Request) {
	var missões models.Missoes
	err := json.NewDecoder(r.Body).Decode(&missões)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
	}

	query := `
	INSERT INTO missoes (nome, descricao, classificacao, rank_ss, rank_s, rank_a, rank_b, rank_c, rank_d, rank_e)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
	`
	var id int

	err = database.Db.QueryRow(query, missões.Nome, missões.Descrição, missões.Classificação, missões.Rank_SS, missões.Rank_S, missões.Rank_A, missões.Rank_B, missões.Rank_C, missões.Rank_D, missões.Rank_E).Scan(id)
	if err != nil {
		http.Error(w, "Erro ao inserir missão", http.StatusInternalServerError)

	}

}
