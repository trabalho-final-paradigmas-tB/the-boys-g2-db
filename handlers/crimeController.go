package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
)

func inserirCrime(w http.ResponseWriter, r *http.Request) {
	query := `INSERT INTO CRIMES (NOME_CRIME, DESCRICAO, DATA_CRIME, HEROI_RESPONSAVEL, SEVERIDADE)
              VALUES ($1, $2, $3, $4, $5)`

	var crime models.Crime
	if err := json.NewDecoder(r.Body).Decode(&crime); err != nil {
		http.Error(w, "Erro ao receber os dados do crime: "+err.Error(), http.StatusBadRequest)
		return
	}

	if crime.NomeCrime == "" || crime.HeroiResponsavel == 0 {
		http.Error(w, "Nome do crime ou herói responsável não fornecido", http.StatusBadRequest)
		return
	}

	if _, err := database.Db.Exec(query,
		crime.NomeCrime,
		crime.Descricao,
		crime.DataCrime,
		crime.HeroiResponsavel,
		crime.Severidade,
	); err != nil {
		http.Error(w, "Erro ao inserir o crime no banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Crime inserido com sucesso!"})
}
