package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	//"github.com/sijms/go-ora/v2"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

// creates a connection and return a connection and an error , connection is returned if success else error if failed
func DbConnect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dbUrl := os.Getenv("DATABASE_URL")
	db, er := sql.Open("postgres", dbUrl)

	if er != nil {
		fmt.Println("Connection Error: ", er)
	}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		return nil, errors.New(err.Error())
	} else {
		log.Println("Successfully Connected")
	}

	return db, nil
}
