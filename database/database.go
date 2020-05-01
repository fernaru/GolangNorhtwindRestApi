package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	connectionString := "root:sasa@tcp(localhost:3307)/northwind"

	databaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error()) //manejo de errores
	}
	return databaseConnection
}
