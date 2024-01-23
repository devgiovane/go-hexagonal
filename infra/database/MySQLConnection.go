package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySQLConnection struct {
	Connection *sql.DB
}

func NewMySQLConnection() *MySQLConnection {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/system")
	if err != nil {
		log.Fatal(err)
	}
	return &MySQLConnection{
		Connection: db,
	}
}

func (m *MySQLConnection) GetConnection() *sql.DB {
	return m.Connection
}
