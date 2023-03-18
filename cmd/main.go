// Tutorial source is https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
	"database/sql"
	"log"
	"os"

	server "github.com/philhanna/restserver"
)

func createDatabaseIfNecessary() {

	checkError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// If the database does not exist, create it
	config, _ := server.NewConfiguration()
	_, err := os.Stat(config.DBNAME)
	if err == nil {
		return
	}
	log.Printf("Creating database %s ...\n", config.DBNAME)

	// Create the empty database
	db, err := sql.Open("sqlite3", config.DBNAME)
	checkError(err)
	defer db.Close() // Make sure it closes

	// Connect, just to check for errors
	err = db.Ping()
	checkError(err)
	log.Printf("Connected to %s\n", config.DBNAME)

	// Load the DDL to create and initialize the table
	dbSql := config.DBSQL
	_, err = db.Exec(dbSql)
	checkError(err)

	// The deferred close will run here:
}

// main is the application mainline
func main() {
	log.Println("Rest API v2.0 - Mux Routers")
	createDatabaseIfNecessary()
	server.HandleRequests()
}
