package handlers

import (
	"backend/database"
	"backend/models"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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
	(CODIGO_HEROI, NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS, DATA_NASCIMENTO)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING CODIGO_HEROI`

	var lastInsertID int
	err := database.Db.QueryRow(insertQuery,
		heroi.ID,
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

func DeletarHerois(w http.ResponseWriter, r *http.Request) {

	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	var heroName string

	if err := json.NewDecoder(r.Body).Decode(&heroName); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := database.Db.Prepare("DELETE FROM heroes WHERE name = ?")
	if err != nil {
		http.Error(w, "Erro ao preparar a consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(heroName)
	if err != nil {
		http.Error(w, "Erro ao deletar herói: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Erro ao verificar linhas afetadas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Herói não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Herói deletado com sucesso"))
}

func BuscarHerois(w http.ResponseWriter, r *http.Request) {
	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	query := "SELECT CODIGO_HEROI, NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS, DATA_NASCIMENTO FROM HEROI"
	rows, err := database.Db.Query(query)
	if err != nil {
		http.Error(w, "Erro ao buscar heróis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []models.Heroi
	for rows.Next() {
		var hero models.Heroi
		if err := rows.Scan(
			&hero.ID,
			&hero.NomeReal,
			&hero.NomeHeroi,
			&hero.Sexo,
			&hero.AlturaHeroi,
			&hero.PesoHeroi,
			&hero.LocalNascimento,
			&hero.Poderes,
			&hero.NivelForca,
			&hero.Popularidade,
			&hero.Status,
			&hero.HistoricoBatalhas,
			&hero.DataNascimento,
		); err != nil {
			http.Error(w, "Erro ao escanear resultados: "+err.Error(), http.StatusInternalServerError)
			return
		}
		heroes = append(heroes, hero)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Erro durante a iteração: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Ta indo os dados")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
	fmt.Println("mandou")
}

func DeletarHeroi(w http.ResponseWriter, r *http.Request) {
	var heroName string

	if err := json.NewDecoder(r.Body).Decode(&heroName); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	query := "SELECT * FROM HEROI WHERE NOME_HEROI ILIKE $1"
	rows, err := database.Db.Query(query, "%"+heroName+"%")
	if err != nil {
		http.Error(w, "Erro ao buscar herói: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []models.Heroi
	for rows.Next() {
		var hero models.Heroi
		if err := rows.Scan(&hero.ID, &hero.NomeReal, &hero.NomeHeroi, &hero.Sexo, &hero.AlturaHeroi, &hero.PesoHeroi, &hero.LocalNascimento, &hero.Poderes, &hero.NivelForca, &hero.Popularidade, &hero.Status, &hero.HistoricoBatalhas, &hero.DataNascimento); err != nil {
			// Handle error here
			fmt.Fprintf(w, "Erro ao escanear resultados: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		heroes = append(heroes, hero)
	}

	if len(heroes) == 0 {
		fmt.Fprintf(w, "Herói não encontrado.\n")
		return
	} else if len(heroes) > 1 {
		fmt.Fprintf(w, "Múltiplos heróis encontrados. Por favor, seja mais específico.\n")
		return
	}

	fmt.Printf("Heroi encontrado: %s\n", heroes[0].NomeHeroi)
	fmt.Print("Deseja deletar este herói? (s/n): ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if strings.TrimSpace(text) == "s" {
		deleteQuery := "DELETE FROM HEROI WHERE CODIGO_HEROI = $1"
		_, err := database.Db.Exec(deleteQuery, heroes[0].ID)
		if err != nil {
			http.Error(w, "Erro ao deletar herói: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Herói deletado com sucesso.\n")
	} else {
		fmt.Fprintf(w, "Deleção cancelada.\n")
	}
}
