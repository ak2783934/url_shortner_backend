package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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
	// fmt.Println(connectionURL)
	// Open a connection to the database
	var err error
	DB, err = sql.Open("mysql", connectionURL)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection is established
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully!")
}

func CloseDBConnections() {
	DB.Close()
}
