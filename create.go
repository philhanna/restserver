package restserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Create adds a new article to the collection, based on JSON
// data contained in the request.  This is a POST method.
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
	rs, err := db.Query(`
	INSERT
	INTO		articles VALUES(null, ?, ?, ?)
	RETURNING 	rowid
	`, article.Title, article.Description, article.Content)
	if err != nil {
		log.Fatal(err)
	}
	defer rs.Close()

	rs.Next()
	var id int
	rs.Scan(&id)
	article.Id = id

	// Respond with the ToJSON string representation of the new article.
	json.NewEncoder(w).Encode(article)
}
