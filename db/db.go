package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTables()
}

func CreateTables() {
	createTable := `CREATE TABLE IF NOT EXISTS links (
  		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	code TEXT NOT NULL,
  		url TEXT NOT NULL,
  		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(createTable)
	if err != nil {
		panic("Could not create links table: " + err.Error())
	}

}
