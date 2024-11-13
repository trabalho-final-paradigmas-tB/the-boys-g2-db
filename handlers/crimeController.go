package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
)

func InserirCrime(w http.ResponseWriter, r *http.Request) {
	query := `INSERT INTO CRIMES (NOME_CRIME, DESCRICAO, DATA_CRIME, HEROI_RESPONSAVEL, SEVERIDADE)
              VALUES ($1, $2, $3, $4, $5) RETURNING ID`

	var crime models.Crime
	if err := json.NewDecoder(r.Body).Decode(&crime); err != nil {
		http.Error(w, "Erro ao receber os dados do crime: "+err.Error(), http.StatusBadRequest)
		return
	}

	if crime.NomeCrime == "" || crime.HeroiResponsavel == 0 {
		http.Error(w, "Nome do crime ou herói responsável não fornecido", http.StatusBadRequest)
		return
	}

	var newID int
	err := database.Db.QueryRow(query,
		crime.NomeCrime,
		crime.Descricao,
		crime.DataCrime,
		crime.HeroiResponsavel,
		crime.Severidade,
	).Scan(&newID)
	if err != nil {
		http.Error(w, "Erro ao inserir o crime no banco de dados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Crime inserido com sucesso!",
		"id":      newID,
	})
}
func ListarCrimes(w http.ResponseWriter, r *http.Request) {
	query := `SELECT ID, NOME_CRIME, DESCRICAO, DATA_CRIME, HEROI_RESPONSAVEL, SEVERIDADE FROM CRIMES`

	rows, err := database.Db.Query(query)
	if err != nil {
		http.Error(w, "Erro ao consultar crimes: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var crimes []models.Crime
	for rows.Next() {
		var crime models.Crime
		if err := rows.Scan(&crime.ID, &crime.NomeCrime, &crime.Descricao, &crime.DataCrime, &crime.HeroiResponsavel, &crime.Severidade); err != nil {
			http.Error(w, "Erro ao ler dados dos crimes: "+err.Error(), http.StatusInternalServerError)
			return
		}
		crimes = append(crimes, crime)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Erro ao iterar sobre os resultados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(crimes)
}
