package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/ak2783934/url_shortner_backend/db"
)

type URLResponse struct {
	LongURL string `json:"long_url"`
	Status  string `json:"status"`
}

func ShortToLong(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the URL parameters
	path := strings.Trim(r.URL.Path, "/")
	// add validations for this path from the URL.
	err := validateShortURL(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you would typically look up the short URL in your database
	// and return the corresponding long URL. For now, we'll just simulate this.

	// find the url in the DB.
	url, err := fetchURLFromURLShortner(path)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	response := URLResponse{
		Status:  "success",
		LongURL: url,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func fetchURLFromURLShortner(path string) (string, error) {
	longURL, err := db.FetchFromDB(path)
	if err != nil {
		log.Println("Error fetching URL from database:", err)
		return "", err
	}
	return longURL, nil
}

func validateShortURL(path string) error {
	// Add your validation logic here
	// For example, check if the path is empty or has invalid characters
	if path == "" {
		return fmt.Errorf("invalid short URL")
	}
	// Add more validations as needed

	var re = regexp.MustCompile(`^[A-Za-z0-9]+$`)

	if !re.MatchString(path) {
		return fmt.Errorf("invalid passed short URL")
	}

	return nil
}
