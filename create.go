package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Create adds a new article to the collection, based on JSON
// data contained in the request.
func Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering Create")

	// Get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)

	// Unmarshal this into a new Article struct
	var article Article
	json.Unmarshal(reqBody, &article)

	// Connect to the database
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Append this to our articles database.
	_, err = db.Exec(`

	INSERT
	INTO	articles VALUES(?, ?, ?, ?)

	`,
		article.Id, article.Title, article.Description, article.Content)
	if err != nil {
		log.Fatal(err)
	}

	// Respond with the ToJSON string representation of the new article.
	json.NewEncoder(w).Encode(article)
}
