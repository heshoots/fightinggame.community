package pkg

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func InitDB() error {
	connStr := os.Getenv("DB_CONN")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
