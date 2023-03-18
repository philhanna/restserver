package restserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var config *Configuration

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Connect opens a connection to the database.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.DBNAME)
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

	var err error
	config, err = NewConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	hostAndPort := fmt.Sprintf("%s:%d", config.HOST, config.PORT)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", GetAll)
	myRouter.HandleFunc("/article", Create).Methods("POST")
	myRouter.HandleFunc("/article/{id}", Delete).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", Update).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", Get)

	log.Printf("Starting server on port %d\n", config.PORT)
	log.Fatal(http.ListenAndServe(hostAndPort, myRouter))
}
