package database

import (
	"database/sql"
	// 
	_ "github.com/go-sql-driver/mysql"
)

//Connect connects to a database and returns teh database to be used 
func Connect() (medicaredb *sql.DB){

	db, err := sql.Open("mysql", "medicareReader:password@tcp(127.0.0.1:3306)/medicare")

	if err != nil {
		panic(err.Error())
	} else {
		medicaredb = db
	}

	return // return medicaredb
}

//ConnectWrite connects to database with writing privileges 
func ConnectWrite() (medicaredb *sql.DB){

	db, err := sql.Open("mysql", "medicareWriter:password@tcp(127.0.0.1:3306)/medicare")

	if err != nil {
		panic(err.Error())
	} else {
		medicaredb = db
	}

	return // return medicaredb
}


