package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ak2783934/url_shortner_backend/db"
	"github.com/ak2783934/url_shortner_backend/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func initEnvVeriables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func registerRoutes() {
	http.HandleFunc("/long-to-short", handlers.LongToShort)
	http.HandleFunc("/", handlers.ShortToLong)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})
}

func main() {
	initEnvVeriables()
	db.InitSQLDB()
	registerRoutes()

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

	defer db.CloseDBConnections()
}
