package database

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDbConnection(connectionString string) *sql.DB {
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error while opening database connection", err)
		os.Exit(1)
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println("Error while opening database connection", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database!")
	return conn
}
