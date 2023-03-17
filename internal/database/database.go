package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DataBase struct {
	db  *sql.DB
	cfg mysql.Config
}

func (database *DataBase) Config(dbName string) {
	database.cfg = mysql.Config{
		User:   os.Getenv("root"),
		Passwd: os.Getenv(""),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: dbName,
	}
}

func (database *DataBase) Connecting() {
	var err error
	database.db, err = sql.Open("mysql", database.cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := database.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
