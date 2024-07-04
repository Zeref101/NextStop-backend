package db

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func ConnectToDB(){

	connection_string := "user=postgres password=postgres dbname=postgres sslmode=disable host=localhost port=5432"


	db, err := sql.Open("postgres", connection_string)

	if err != nil {
		log.Fatal(err)
	}

	//  verifying the connection by pinging the db

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

}

func CloseDB(){

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}