package db

import (
	"fmt"
	"log"
)

func SaveToDB(longURL string, shortURL string) error {
	// Prepare the SQL statement
	result, err := DB.Exec("INSERT INTO url_shortner (long_url, short_url) VALUES (?, ?)", longURL, shortURL)
	if err != nil {
		return err
	}
	// Get the number of affected rows
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rowsAffected)
	return nil
}

func FetchFromDB(shortURL string) (string, error) {
	// Prepare the SQL statement
	rows, err := DB.Query("SELECT long_url FROM url_shortner WHERE short_url = ?", shortURL)
	if err != nil {
		log.Println("Error executing query:", err)
		return "", err
	}
	defer rows.Close()
	var longURL string
	if rows.Next() {
		err := rows.Scan(&longURL)
		if err != nil {
			log.Println("Error scanning row:", err)
			return "", err
		}
	}

	return longURL, nil
}
