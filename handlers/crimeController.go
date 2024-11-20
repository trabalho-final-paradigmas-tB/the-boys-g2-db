package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func InserirCrime(w http.ResponseWriter, r *http.Request) {
	query := `INSERT INTO CRIMES (NOME, DESCRICAO, DATA, HEROI_RESPONSAVEL, SEVERIDADE)
              VALUES ($1, $2, $3, $4, $5) RETURNING ID`

	var crime models.Crime
	if err := json.NewDecoder(r.Body).Decode(&crime); err != nil {
		http.Error(w, "Erro ao receber os dados do crime: "+err.Error(), http.StatusBadRequest)
		return
	}

	if crime.NomeCrime == "" || crime.HeroiResponsavel == "" {
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

	var ajustePopularidade int
	switch crime.Severidade {
	case "leve":
		ajustePopularidade = -1
	case "moderada":
		ajustePopularidade = -3
	case "grave":
		ajustePopularidade = -5
	default:
		ajustePopularidade = -2
	}

	updateQuery := `UPDATE HEROI SET POPULARIDADE = POPULARIDADE + $1 WHERE CODIGO_HEROI = $2`
	_, err = database.Db.Exec(updateQuery, ajustePopularidade, crime.HeroiResponsavel)
	if err != nil {
		http.Error(w, "Erro ao atualizar a popularidade do herói: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Crime inserido com sucesso!",
		"id":      newID,
	})
}

func ListarCrimes(w http.ResponseWriter, r *http.Request) {
	query := `SELECT ID, NOME, DESCRICAO, DATA, HEROI_RESPONSAVEL, SEVERIDADE 
              FROM CRIMES WHERE OCULTO = false`

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

func OcultarCrime(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var count int
	err := database.Db.QueryRow(`SELECT COUNT(*) FROM CRIMES WHERE ID = $1`, id).Scan(&count)
	if err != nil {
		http.Error(w, "Erro ao verificar se o crime existe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Crime não encontrado.", http.StatusNotFound)
		return
	}

	query := `UPDATE CRIMES SET OCULTO = true WHERE ID = $1`
	_, err = database.Db.Exec(query, id)
	if err != nil {
		http.Error(w, "Erro ao ocultar o crime: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Crime ocultado com sucesso!",
	})
}
