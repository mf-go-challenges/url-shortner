package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func InitDB() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env file found")
		}
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTables()
}

func CreateTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
  		id SERIAL PRIMARY KEY,
    	username TEXT NOT NULL,
  		password TEXT NOT NULL,
  		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}

	createLinksTable := `CREATE TABLE IF NOT EXISTS links (
  		id SERIAL PRIMARY KEY,
    	code TEXT NOT NULL,
  		url TEXT NOT NULL,
  		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	    user_id INTEGER ,
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createLinksTable)
	if err != nil {
		panic("Could not create links table: " + err.Error())
	}
}
