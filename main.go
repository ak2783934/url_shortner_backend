package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ak2783934/url_shortner_backend/db"
	"github.com/ak2783934/url_shortner_backend/handlers"
	"github.com/ak2783934/url_shortner_backend/middleware"
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
	http.HandleFunc("/long-to-short", middleware.EnableCORS(handlers.LongToShort))
	http.HandleFunc("/", middleware.EnableCORS(handlers.ShortToLong))
	http.HandleFunc("/health", middleware.EnableCORS(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	}))
}

func main() {
	initEnvVeriables()
	db.InitSQLDB()
	registerRoutes()

	port := os.Getenv("APP_PORT")
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

	defer db.CloseDBConnections()
}
