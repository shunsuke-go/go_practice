package main

import (
	"go_practice/infrastracture/db"
	"go_practice/interfaceAdapter/http/routers"
)

func main() {
	dbconn := db.NewDBConnector()
	psqlConn := dbconn.PostgresConnector()
	routers.Run(psqlConn.Conn)
}
