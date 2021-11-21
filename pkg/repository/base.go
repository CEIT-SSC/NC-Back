package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgres golang driver
	"os"
)


func CreateConnection() (*sql.DB, error) {
	//Open the connection
	//err := godotenv.Load(".env")
	//
	//if err != nil {
	//	log.Fatalf("Error loading .env file")
	//}

	dbUsername := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbUri := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", dbHost,dbPort,  dbUsername,password, dbName )
	fmt.Println(dbUri)

	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return db, nil
}
