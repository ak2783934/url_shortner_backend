package handlers

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
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

func generateAndSaveShortURL(longURL string) (string, error) {
	var shortLink string
	for {
		result := make([]byte, numCharsShortLink)
		for i := 0; i < numCharsShortLink; i++ {
			index := seededRand.Intn(len(alphabet))
			result[i] = alphabet[index]
		}
		shortLink = string(result)

		// check if this link is already there in db?
		url, err := db.FetchFromDB(shortLink)
		if err != nil || url == "" {
			break
		}
	}

	db.SaveToDB(longURL, shortLink)
	return "https://url_shorter.com/" + shortLink, nil
}
