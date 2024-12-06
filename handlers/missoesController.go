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

	// Validação de dificuldade
	if missoes.Dificuldade < 1 || missoes.Dificuldade > 10 {
		http.Error(w, "Dificuldade inválida. Deve estar entre 1 e 10.", http.StatusBadRequest)
		return
	}

	// Validação de recompensa
	if missoes.RecompensaValor <= 0 {
		http.Error(w, "O valor da recompensa deve ser maior que 0.", http.StatusBadRequest)
		return
	}

	query := `
    INSERT INTO MISSOES (NOME, DESCRICAO, CLASSIFICACAO, DIFICULDADE, HEROIS, RECOMPENSA_TIPO, RECOMPENSA_VALOR)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING ID
    `
	var id int

	err = database.Db.QueryRow(query,
		missoes.Nome,
		missoes.Descrição,
		missoes.Classificação,
		missoes.Dificuldade,
		pq.Array(missoes.Herois),
		missoes.RecompensaTipo,
		missoes.RecompensaValor,
	).Scan(&id)

	if err != nil {
		log.Printf("Erro ao inserir missão no banco de dados: %v", err)
		http.Error(w, "Erro ao inserir missão: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurando a resposta
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"mensagem": "Missão inserida com sucesso",
		"status":   "sucesso",
		"missao":   missoes.Nome,
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

	query := "SELECT id, nome, descricao, classificacao, dificuldade, herois, recompensa_tipo, recompensa_valor FROM MISSOES"
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
		err := rows.Scan(&missao.ID, &missao.Nome, &missao.Descrição, &missao.Classificação, &missao.Dificuldade, pq.Array(&missao.Herois), &missao.RecompensaTipo, &missao.RecompensaValor)
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
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}

	// Decodifica os dados da missão do corpo da requisição
	var missao models.Missoes
	err := json.NewDecoder(r.Body).Decode(&missao)
	if err != nil {
		http.Error(w, "Erro ao ler dados da missão: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validação da dificuldade da missão
	if missao.Dificuldade < 1 || missao.Dificuldade > 10 {
		http.Error(w, "Nível de dificuldade inválido. Deve estar entre 1 e 10.", http.StatusBadRequest)
		return
	}

	var resultados []map[string]interface{}

	// Processa cada herói da missão
	for _, nomeHeroi := range missao.Herois {
		var heroi models.Heroi

		// Busca as informações do herói pelo nome no banco de dados
		err = database.Db.QueryRow(
			"SELECT NOME_HEROI, NIVEL_FORCA, POPULARIDADE, CODIGO_HEROI FROM HEROI WHERE NOME_HEROI = $1",
			nomeHeroi,
		).Scan(&heroi.NomeHeroi, &heroi.NivelForca, &heroi.Popularidade, &heroi.CodigoHeroi)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("Herói não encontrado: %s", nomeHeroi)
				resultados = append(resultados, map[string]interface{}{
					"mensagem": "Herói não encontrado",
					"heroi":    nomeHeroi,
				})
				continue
			}

			http.Error(w, "Erro ao buscar os dados do herói "+nomeHeroi+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Validação do nível de força do herói
		if heroi.NivelForca < 0 || heroi.NivelForca > 100 {
			http.Error(w, "Nível de força inválido para o herói "+heroi.NomeHeroi, http.StatusBadRequest)
			return
		}

		// Determinação do sucesso da missão para o herói
		sucesso := missao.Dificuldade <= heroi.NivelForca
		if !sucesso {
			resultados = append(resultados, map[string]interface{}{
				"mensagem": "Missão foi um fracasso para o herói",
				"heroi":    heroi.NomeHeroi,
			})
		}

		// Cálculo da recompensa
		var valorRecompensa int
		if missao.RecompensaValor > 0 {
			valorRecompensa = missao.RecompensaValor
		} else if missao.Dificuldade > 7 {
			valorRecompensa = 10
		} else {
			valorRecompensa = 5
		}

		// Aplicação da recompensa
		if missao.RecompensaTipo == "Força" {
			heroi.NivelForca += valorRecompensa
			if heroi.NivelForca > 100 {
				heroi.NivelForca = 100
			}
			_, err = database.Db.Exec("UPDATE HEROI SET NIVEL_FORCA = $1 WHERE CODIGO_HEROI = $2", heroi.NivelForca, heroi.CodigoHeroi)
			if err != nil {
				http.Error(w, "Erro ao atualizar nível de força do herói "+heroi.NomeHeroi+": "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else if missao.RecompensaTipo == "Popularidade" {
			heroi.Popularidade += valorRecompensa
			if heroi.Popularidade > 100 {
				heroi.Popularidade = 100
			}
			_, err = database.Db.Exec("UPDATE HEROI SET POPULARIDADE = $1 WHERE CODIGO_HEROI = $2", heroi.Popularidade, heroi.CodigoHeroi)
			if err != nil {
				http.Error(w, "Erro ao atualizar popularidade do herói "+heroi.NomeHeroi+": "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if err != nil {
			http.Error(w, "Erro ao atualizar o herói "+heroi.NomeHeroi+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Adiciona resultado do herói à lista
		resultados = append(resultados, map[string]interface{}{
			"mensagem":          "Missão concluída com sucesso para o herói",
			"heroi":             heroi.NomeHeroi,
			"novo_nivel_forca":  heroi.NivelForca,
			"nova_popularidade": heroi.Popularidade,
		})
	}

	// Retorna a resposta com os resultados da missão
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"resultados": resultados,
		"missao":     missao.Nome,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao enviar a resposta: "+err.Error(), http.StatusInternalServerError)
	}
}
