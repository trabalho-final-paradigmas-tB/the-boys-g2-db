package handlers

import (
		"net/http"
 		"database/sql"
 	    "fmt"
		)

func InserirHeroi(w http.ResponseWriter, r *http.Request) {
    dsn := "about_hero.sql"

    db, err := sql.Open("sqlite3", dsn)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Conectado ao banco de dados com sucesso!")

    // Exemplo de inserção de dados
    insertQuery := "INSERT INTO sua_tabela (coluna1, coluna2) VALUES (?, ?)"
    res, err := db.Exec(insertQuery, "valor1", "valor2")
    if err != nil {
        panic(err)
    }

    lastInsertID, err := res.LastInsertId()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Último ID inserido: %d\n", lastInsertID)

    // Exemplo de atualização de dados
    updateQuery := "UPDATE sua_tabela SET coluna1 = ? WHERE id = ?"
    res, err = db.Exec(updateQuery, "novo_valor", lastInsertID)
    if err != nil {
        panic(err)
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Número de linhas afetadas: %d\n", rowsAffected)

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

	// quando for fazer o db definido a gente mexe nisso
}


