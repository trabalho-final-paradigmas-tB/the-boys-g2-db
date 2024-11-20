package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"strings"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func InserirMissao(w http.ResponseWriter, r *http.Request) {
	if database.Db == nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}

	var missoes models.Missoes
	err := json.NewDecoder(r.Body).Decode(&missoes)
	if err != nil {
		http.Error(w, "Erro ao ler dados: "+err.Error(), http.StatusBadRequest)
		return
	}

	if missoes.Dificuldade < 1 || missoes.Dificuldade > 10 {
		http.Error(w, "Dificuldade inválida. Deve estar entre 1 e 10.", http.StatusBadRequest)
		return
	}

	query := `
    INSERT INTO MISSOES (NOME, DESCRICAO, CLASSIFICACAO, DIFICULDADE, HEROIS)
        VALUES ($1, $2, $3, $4, $5) RETURNING ID
    `
	var id int

	err = database.Db.QueryRow(query,
		missoes.Nome,
		missoes.Descrição,
		missoes.Classificação,
		missoes.Dificuldade,
		pq.Array(missoes.Herois),
	).Scan(&id)

	if err != nil {
		log.Printf("Erro ao inserir missão no banco de dados: %v", err)
		http.Error(w, "Erro ao inserir missão: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"mensagem":              "Missão inserida com sucesso",
		"status":                "sucesso",
		"missao":                missoes.Nome,
		"descrição":             missoes.Descrição,
		"dificuldade_escolhida": missoes.Dificuldade,
		"herois_na_missao":      missoes.Herois,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao enviar a resposta: "+err.Error(), http.StatusInternalServerError)
	}
}

func ListadeMissões(w http.ResponseWriter, r *http.Request) {
	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	query := "SELECT id, nome, descricao, classificacao, dificuldade, herois FROM MISSOES"
	rows, err := database.Db.Query(query)
	if err != nil {
		http.Error(w, "Erro ao listar missões: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao executar query: %v", err)
		return
	}
	defer rows.Close()

	var missoes []models.Missoes

	for rows.Next() {
		var missao models.Missoes
		err := rows.Scan(&missao.ID, &missao.Nome, &missao.Descrição, &missao.Classificação, &missao.Dificuldade, pq.Array(&missao.Herois))
		if err != nil {
			http.Error(w, "Erro ao escanear missão: "+err.Error(), http.StatusInternalServerError)
			log.Printf("Erro ao escanear linha: %v", err)
			return
		}

		missoes = append(missoes, missao)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro na leitura das linhas: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro no iterador de linhas: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(missoes); err != nil {
		http.Error(w, "Erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao codificar JSON: %v", err)
	}
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
	query := "SELECT NOME FROM MISSOES WHERE id = $1"
	err = database.Db.QueryRow(query, missaoID).Scan(&nomeMissao)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Missão não encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao buscar missão: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	deleteQuery := "DELETE FROM MISSOES WHERE id = $1"
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

	query := `UPDATE MISSOES
        SET NOME = $1, DESCRICAO = $2, CLASSIFICACAO = $3, DIFICULDADE = $4
        WHERE id = $5`

	if missao.Dificuldade < 1 || missao.Dificuldade > 10 {
		http.Error(w, "Dificuldade Inválida", http.StatusBadRequest)
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

func Resultadomissão(w http.ResponseWriter, r *http.Request) {
	if database.Db == nil {
		http.Error(w, "Erro ao conectar banco de dados", http.StatusInternalServerError)
		return
	}
	var missao models.Missoes
	var herois models.Heroi

	err := json.NewDecoder(r.Body).Decode(&missao)
	if err != nil {
		http.Error(w, "Erro ao ler dados da missão", http.StatusBadRequest)
		return
	}

	if missao.Dificuldade < 1 || missao.Dificuldade > 10 {
		http.Error(w, "Nivel de dificuldade invalido", http.StatusBadRequest)
		return
	}
	if herois.NivelForca < 1 || herois.NivelForca > 10 {
		http.Error(w, "Nivel de força invalido", http.StatusBadRequest)
		return
	}

	if missao.Dificuldade > herois.NivelForca {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"mensagem": "Missão foi um Fracasso",
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Erro ao enviar a resposta", http.StatusInternalServerError)
		}
		return
	}

	if missao.Dificuldade <= herois.NivelForca {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"mensagem": "Missão concluída com sucesso",
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Erro ao enviar a resposta", http.StatusInternalServerError)
		}
	}
}
