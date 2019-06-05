package db

import (
	"database/sql"

	//Importado para conectar com postgresql
	_ "github.com/lib/pq"
)

// ObtenhaConexao obtém uma instância de conexão com o banco de dados.
func ObtenhaConexao() (*sql.DB, error) {
	connStr := "user=postgres dbname=uatidb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
