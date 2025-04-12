package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var DB *sql.DB

func InitSQLDB() {
	// Initialize the database connection
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASS")
	mysqlHost := os.Getenv("MYSQL_HOST_GO")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDBNAME := os.Getenv("MYSQL_DB_NAME")
	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDBNAME)

	// Open a connection to the database
	var err error

	for i := 0; i < 10; i++ {
		DB, err = sql.Open("mysql", connectionURL)
		if err == nil {
			err = DB.Ping()
			if err == nil {
				break
			}
		}
		log.Println("Waiting for MySQL to be ready...")
		time.Sleep(3 * time.Second)
	}

	fmt.Println("Database connected successfully!")
}

func CloseDBConnections() {
	DB.Close()
}
