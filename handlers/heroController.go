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
	"github.com/lib/pq"
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

	if heroi.HistoricoBatalhas == nil {
		heroi.HistoricoBatalhas = []int{0, 0}
	}

	insertQuery := `INSERT INTO HEROI
	(NOME_REAL, NOME_HEROI, SEXO, ALTURA, PESO, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS, DATA_NASCIMENTO)
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
		pq.Array(heroi.HistoricoBatalhas),
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

	rows, err := database.Db.Query(`
		SELECT 
			CODIGO_HEROI, NOME_REAL, NOME_HEROI, SEXO, ALTURA, PESO, DATA_NASCIMENTO, 
			LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS 
		FROM HEROI
	`)
	if err != nil {
		http.Error(w, "Erro ao listar heróis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var herois []models.Heroi

	for rows.Next() {
		var heroi models.Heroi
		err := rows.Scan(
			&heroi.CodigoHeroi,
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
			&heroi.HistoricoBatalhas, // Usa o tipo IntArray
		)
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

	query := `SELECT * FROM HEROI WHERE CODIGO_HEROI = $1`

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
		pq.Array(heroi.HistoricoBatalhas),
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

func ListarHeroisPorNome(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")

	if nome == "" {
		http.Error(w, "Nome do herói é obrigatório", http.StatusBadRequest)
		return
	}

	query := `SELECT * FROM HEROI WHERE NOME_HEROI ILIKE $1`

	rows, err := database.Db.Query(query, "%"+nome+"%")
	if err != nil {
		http.Error(w, "Erro ao consultar heróis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var herois []*models.Heroi
	for rows.Next() {
		var heroi models.Heroi
		err := rows.Scan(
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
			http.Error(w, "Erro ao ler dados do herói: "+err.Error(), http.StatusInternalServerError)
			return
		}
		herois = append(herois, &heroi)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro ao iterar sobre os resultados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(herois)
}

func ListarHeroisPorStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		http.Error(w, "Status é obrigatório", http.StatusBadRequest)
		return
	}

	query := `SELECT * FROM HEROI WHERE STATUS = $1`

	rows, err := database.Db.Query(query, status)
	if err != nil {
		http.Error(w, "Erro ao consultar heróis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var herois []models.Heroi
	for rows.Next() {
		var heroi models.Heroi
		err := rows.Scan(
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
			&heroi.HistoricoBatalhas)
		if err != nil {
			http.Error(w, "Erro ao escanear herói: "+err.Error(), http.StatusInternalServerError)
			return
		}
		herois = append(herois, heroi)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Erro ao iterar sobre os resultados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(herois)
}

func ListarHeroisPorPolularidade(w http.ResponseWriter, r *http.Request) {
	popularidadestr := r.URL.Query().Get("popularidade")
	popula, err := strconv.Atoi(popularidadestr)
	if err != nil {
		http.Error(w, "Imput inválido", http.StatusBadRequest)
		return
	}

	query := `SELECT *  FROM HEROI WHERE POPULARIDE = $1`

	var heroi models.Heroi
	err = database.Db.QueryRow(query, popula).Scan(
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
			http.Error(w, "Erro ao consultar herói: "+err.Error()+". Consulta: "+query, http.StatusInternalServerError)
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
	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	var heroi models.Heroi
	if err := json.NewDecoder(r.Body).Decode(&heroi); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	heroiIDStr := vars["id"]
	heroiID, err := strconv.Atoi(heroiIDStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	var heroexits bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM HEROI WHERE CODIGO_HEROI = $1)"
	err = database.Db.QueryRow(checkQuery, heroiID).Scan(&heroexits)
	if err != nil || !heroexits {
		http.Error(w, "Heroi não encontrado", http.StatusNotFound)
		return
	}

	query := `UPDATE HEROI
        SET NOME_REAL = $1, NOME_HEROI = $2, SEXO = $3, ALTURA = $4, PESO = $5, LOCAL_NASCIMENTO = $6, PODERES = $7, NIVEL_FORCA = $8, POPULARIDADE = $9, STATUS = $10, HISTORICO_BATALHAS = $11, DATA_NASCIMENTO = $12
        WHERE CODIGO_HEROI = $13`

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
		pq.Array(heroi.HistoricoBatalhas),
		heroi.DataNascimento,
		heroiID,
	)
	if err != nil {
		http.Error(w, "Erro ao atualizar herói: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Herói atualizado com sucesso!"})
}
