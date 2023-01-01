package db

import (
	"database/sql"
	"fmt"
	"go_practice/config"
	"go_practice/infrastracture/db/postgres"
)
type DB struct {
	Conn *sql.DB
}

type DBConnector struct {
}

type IDBConnector interface {
	PostgresConnector() (*DB)
}

func NewDBConnector() IDBConnector {
	return &DBConnector{}
}

func (dbconn *DBConnector) PostgresConnector() *DB {
	psqlConf := config.GetPostgresConfig()
	conn, err := sql.Open("postgres", postgres.ConfigToString(psqlConf))
	if err != nil {
		fmt.Print("connect error: ", err)
	}
	psqlConnector := &DB{
		Conn: conn,
	}
	return psqlConnector
}