package postgres

import (
	"database/sql"
	"fmt"
)

type PostgresConnector struct {
	Conn *sql.DB
}

func NewPostgresConnector () *PostgresConnector {
	conn, err := sql.Open("postgres", "host=host.docker.internal port=5432 user=postgres password=password dbname=postgres sslmode=disable search_path=go_practice")
	if err != nil {
		fmt.Print("connect error: ", err)
	}
	psqlConnector := &PostgresConnector{
		Conn: conn,
	}
	return psqlConnector
}