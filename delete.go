package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Delete gets an article number from the request and deletes the
// corresponding article in the collection.
func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering DeleteArticle")

	// Parse the request parameters and identify the requested ID
	vars := mux.Vars(r)
	id := vars["id"]

	// Connect to the database
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the article
	rs, err := db.Query(`
	
	DELETE
	FROM	articles
	WHERE	id=?
	
	`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rs.Close()
}
