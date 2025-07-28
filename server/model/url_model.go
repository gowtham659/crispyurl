package model

import (
	"fmt"
	"log"
	"urlshortnerreactgo/database"
)

var con, dber = database.DbConnect()

func InsertUrl(original_url string, baseUrl string, shot_code string) bool {
	if dber != nil {
		log.Fatalln("connection error")
	}

	query, er := con.Exec("INSERT INTO url_shortener (original_url,shorted_url,short_code) VALUES ($1, $2, $3)", original_url, (baseUrl + shot_code), shot_code)
	if er != nil {
		fmt.Println("insert query error", er)
	}
	if r, er := query.RowsAffected(); er == nil && r > 0 {
		return true
	}
	return false
}

// GetOriginalURL retrieves the original URL using the short code
func GetOriginalURL(shortCode string) (string, error) {
	if dber != nil {
		log.Fatalln("connection error")
	}
	var originalURL string
	query := "SELECT original_url FROM url_shortener WHERE short_code = $1"
	err := con.QueryRow(query, shortCode).Scan(&originalURL)
	if err != nil {
		return "", err //errors.New("no matching rows")
	}
	return originalURL, nil
}
