package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBD() *sql.DB {
	conexao := "user=postgres dbname=Alura_Loja password=abc123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
