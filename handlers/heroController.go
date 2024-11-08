package handlers

import (
	"backend/database"
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
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
	(NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS, DATA_NASCIMENTO)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING CODIGO_HEROI`

	var lastInsertID int
	err := database.Db.QueryRow(insertQuery,
		heroi.CodigoHeroi,
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

func ListarHerois(w http.ResponseWriter, r *http.Request) {

	if database.Db == nil {
		http.Error(w, "Erro de conexão com o banco de dados", http.StatusInternalServerError)
		return
	}

	var heroi models.Heroi
	if err := json.NewDecoder(r.Body).Decode(&heroi); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := database.Db.Query("SHOW TABLES") // Coleta da tabela de HEROIS e suas caracteristicas

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var tableName string // Imprimir tabela abaixo vvvv

	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(tableName)
	}
}

func ListarHeroiPorID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	query := `SELECT NOME_REAL, NOME_HEROI, SEXO, ALTURA_HEROI, PESO_HEROI, DATA_NASCIMENTO,
	LOCAL_NASCIMENTO, PODERES, NIVEL_FORCA, POPULARIDADE, STATUS, HISTORICO_BATALHAS
	FROM HEROI WHERE CODIGO_HEROI = $1`

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
		&heroi.HistoricoBatalhas,
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

func DeletarHeroi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	heroiIDStr := vars["id"]

	heroid, err := strconv.Atoi(heroiIDStr)
	if err != nil {
		http.Error(w, "ID inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// consultar o herói antes de deletá-lo para obter o nome

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

	// executar o comando DELETE

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

	// retornar o ID e o nome do herói que foi deletado em vez só de 'herois'

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

	var input string
	fmt.Print("Digite o nome do herói (real ou de herói) ou o código para modificar: ")
	fmt.Scanln(&input)

	// Consulta ao banco de dados
	err := database.Db.QueryRow("SELECT * FROM herois WHERE CODIGO_HEROI=? OR NOME_REAL=? OR NOME_HEROI=?", input, input, input).Scan(&heroi.CodigoHeroi, &heroi.NomeReal, &heroi.NomeHeroi)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Herói não encontrado.")
		} else {
			fmt.Println("Erro ao buscar herói:", err)
		}
	}

	// Exibir informações do herói
	fmt.Printf("Herói encontrado: \n Codigo heroi %s \nNome real: %s\nNome de herói: %s\n", heroi.CodigoHeroi, heroi.NomeReal, heroi.NomeHeroi)

	// Oferecer opções de modificação
	fmt.Println("Digite o atributo a ser modificado (ou 0 para cancelar):")
	var atributo string
	fmt.Scanln(&atributo)

	fmt.Println("Digite o novo valor (número inteiro):")
	var novoValor int
	fmt.Scanln(&novoValor)

	// Obter o tipo do campo a partir da estrutura do herói
	v := reflect.ValueOf(heroi).Elem()
	field := v.FieldByName(atributo)
	if !field.IsValid() {
		fmt.Println("Atributo inválido")
		return 0
	}

	// Construir a query de forma segura usando parâmetros
	query := "UPDATE herois SET ? = ? WHERE id = ?"

	// Executar a query com os valores parametrizados
	_, err = database.Db.Exec(query, atributo, novoValor, heroi.CodigoHeroi)
	if err != nil {
		fmt.Println("Erro ao atualizar herói:", err)
		return err
	}

	// Atualizar o objeto herói em memória
	field.SetInt(int64(novoValor))

	fmt.Println("Atributo atualizado com sucesso!")
	return nil
}
