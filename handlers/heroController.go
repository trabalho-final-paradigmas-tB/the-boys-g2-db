package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Usuario struct {
	ID   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func InserirUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	insertQuery := `INSERT INTO USER (ID, AGE, NAME) VALUES (@ID, @AGE, @NAME)`

	res, err := database.DB.Exec(insertQuery,
		sql.Named("ID", usuario.ID),
		sql.Named("AGE", usuario.Age),
		sql.Named("NAME", usuario.Name),
	)

	if err != nil {
		http.Error(w, "Erro ao inserir usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	lastInsertID, _ := res.LastInsertId()
	fmt.Fprintf(w, "Usuário inserido com sucesso! Último ID inserido: %d\n", lastInsertID)
}

func InserirHeroi(w http.ResponseWriter, r *http.Request) {
	var heroi models.Heroi
	if err := json.NewDecoder(r.Body).Decode(&heroi); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	insertQuery := `INSERT INTO HEROI
    (ID, NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS, DATA_NASCIMENTO)
    VALUES (@ID, @NOME_REAL, @NOME_HEROI, @SEXO, @ALTURA_HEROI, @PESO_HEROI, @LOCAL_NASCIMENTO, @PODERES, @NIVEL_FORCA, @POPULARIDADE, @STATUS, @HISTORICO_BATALHAS, @DATA_NASCIMENTO)`

	res, err := database.DB.Exec(insertQuery,
		sql.Named("ID", heroi.ID),
		sql.Named("NOME_REAL", heroi.NomeReal),
		sql.Named("NOME_HEROI", heroi.NomeHeroi),
		sql.Named("SEXO", heroi.Sexo),
		sql.Named("ALTURA_HEROI", heroi.AlturaHeroi),
		sql.Named("PESO_HEROI", heroi.PesoHeroi),
		sql.Named("LOCAL_NASCIMENTO", heroi.LocalNascimento),
		sql.Named("PODERES", heroi.Poderes),
		sql.Named("NIVEL_FORCA", heroi.NivelForca),
		sql.Named("POPULARIDADE", heroi.Popularidade),
		sql.Named("STATUS", heroi.Status),
		sql.Named("HISTORICO_BATALHAS", heroi.HistoricoBatalhas),
		sql.Named("DATA_NASCIMENTO", heroi.DataNascimento),
	)

	if err != nil {
		http.Error(w, "Erro ao inserir herói: "+err.Error(), http.StatusInternalServerError)
		return
	}

	lastInsertID, _ := res.LastInsertId()
	fmt.Fprintf(w, "Herói inserido com sucesso! Último ID inserido: %d\n", lastInsertID)
}

func ListarHerois(w http.ResponseWriter, r *http.Request) {
	query := "SELECT NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, DATA_NASCIMENTO, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS FROM herois"

	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Erro ao consultar heróis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var herois []models.Heroi
	for rows.Next() {
		var heroi models.Heroi
		if err := rows.Scan(&heroi.NomeReal, &heroi.NomeHeroi, &heroi.Sexo, &heroi.AlturaHeroi, &heroi.PesoHeroi, &heroi.DataNascimento, &heroi.LocalNascimento, &heroi.Poderes, &heroi.NivelForca, &heroi.Popularidade, &heroi.Status, &heroi.HistoricoBatalhas); err != nil {
			http.Error(w, "Erro ao escanear herói: "+err.Error(), http.StatusInternalServerError)
			return
		}
		herois = append(herois, heroi)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(herois)
}

/*
func AtualizarHeroi(w http.ResponseWriter, r *http.Request) {

	// Exemplo de atualização de dados
	updateQuery := "UPDATE sua_tabela SET coluna1 = ? WHERE id = ?"
	w, err = db.Exec(updateQuery, "novo_valor", lastInsertID)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Número de linhas afetadas: %d\n", rowsAffected)
}

func DeletarHeroi(w http.ResponseWriter, r *http.Request) {

	// Exemplo de exclusão de dados
	deleteQuery := "DELETE FROM sua_tabela WHERE id = ?"
	res, err = db.Exec(deleteQuery, lastInsertID)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Número de linhas excluídas: %d\n", rowsAffected)
}
*/
