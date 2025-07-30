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
)

// creates a connection and return a connection and an error , connection is returned if success else error if failed
func DbConnect() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	db, er := sql.Open("postgres", dbUrl)
	//db,er := sqlx.Connect("postgres","user=postgres dbname=urlshortnerdb sslmode=disable password=gowtham123 host=localhost")
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
