package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnDB() *sql.DB {
	dsn := os.Getenv("MySQL_URL")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
