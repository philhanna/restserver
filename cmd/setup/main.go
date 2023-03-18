package main

import (
	"database/sql"
	"log"
	server "github.com/philhanna/restserver"
	_ "github.com/mattn/go-sqlite3"
)

var config server.Configuration

func main() {

	check := func (err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// Load the configuration
	config, err := server.NewConfiguration()
	check(err)

	// Create the empty database
	db, err := sql.Open("sqlite3", config.DBNAME)
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

	// Load the DDL to create and initialize the table
	dbSql := config.DBSQL
	_, err = db.Exec(dbSql)
	check(err)

}