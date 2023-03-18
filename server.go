package webserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// ---------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------

const (
	HOST   = "localhost"
	PORT   = 10000
	DBNAME = "articles.db"
)

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Connect opens a connection to the database.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBNAME)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// HandleRequests registers all the routers and starts the server.
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", GetAll)
	myRouter.HandleFunc("/article", Create).Methods("POST")
	myRouter.HandleFunc("/article/{id}", Delete).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", Update).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", Get)
	
	log.Printf("Starting server on port %d\n", PORT)
	hostAndPort := fmt.Sprintf("%s:%d", HOST, PORT)
	log.Fatal(http.ListenAndServe(hostAndPort, myRouter))
}
