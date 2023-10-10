package utilities

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var cfg = mysql.Config{
	User: os.Getenv("DBUSER"),
}

func connect() {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}
}
