package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
	"log"
	"os"
)

var CreateConnection = func() *sql.DB {
	// Open the connection
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
