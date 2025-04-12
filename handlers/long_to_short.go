package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/ak2783934/url_shortner_backend/db"
)

type LongToShortRequest struct {
	LongURL string `json:"long_url"`
}

type LongToShortResponse struct {
	ShortURL string `json:"short_url"`
	Status   string `json:"status"`
}

const (
	numCharsShortLink = 7
	alphabet          = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func LongToShort(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var reqBody LongToShortRequest
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
	}

	err = validateLongUrl(reqBody.LongURL)
	if err != nil {
		http.Error(w, "Invalid long url link", http.StatusBadRequest)
	}

	shortURL, err := generateAndSaveShortURL(reqBody.LongURL)
	if err != nil {
		http.Error(w, "Error generating short url", http.StatusInternalServerError)
	}

	response := LongToShortResponse{
		Status:   "success",
		ShortURL: shortURL,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func validateLongUrl(longURL string) error {
	if longURL == "" {
		return fmt.Errorf("invalid short URL")
	}
	// Add more validations as needed
	re := regexp.MustCompile(`^(http|https):\/\/[^\s/$.?#].[^\s]*$`)
	if !re.MatchString(longURL) {
		return fmt.Errorf("invalid URL format")
	}
	return nil
}

func generateAndSaveShortURL(longURL string) (string, error) {
	// first check if this long url is already present in the db directly?
	// if yes, then return that value itself
	shortURL, _ := db.FetchShorURLFromDB(longURL)
	if shortURL != "" {
		return shortURL, nil
	}

	var shortLink string
	for {
		result := make([]byte, numCharsShortLink)
		for i := 0; i < numCharsShortLink; i++ {
			index := seededRand.Intn(len(alphabet))
			result[i] = alphabet[index]
		}
		shortLink = string(result)

		// check if this link is already there in db?
		url, err := db.FetchLongURLFromDB(shortLink)
		if err != nil || url == "" {
			break
		}
	}

	saveErr := db.SaveToDB(longURL, shortLink)
	if saveErr != nil {
		return "", saveErr
	}
	fmt.Println("succesfully saved long url ", longURL, " short url", shortLink)
	return shortLink, nil
}
