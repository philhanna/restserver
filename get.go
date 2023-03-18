package webserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Get gets an article number from the request and
// returns the corresponding article in the collection.
func Get(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering Get")

	// Get the requested Id
	vars := mux.Vars(r)
	key := vars["id"]
	log.Printf("Getting article %q\n", key)

	// Connect to the database
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Select the requested article
	rs, err := db.Query(`

		SELECT	id,
				title,
				description,
				content
		FROM	articles
		WHERE	Id=?

		`, key)
	if err != nil {
		log.Fatal(err)
	}
	defer rs.Close()

	// Read the result set into a new Article structure
	article := Article{}
	if rs.Next() {
		rs.Scan(&article.Id, &article.Title, &article.Description, &article.Content)

		// If it is found, return its JSON representation
		json.NewEncoder(w).Encode(article)
	} else {
		// TODO test this path
		// Otherwise, return a 404 eror
		http.Error(w, "Article not found", http.StatusNotFound)
	}
}
