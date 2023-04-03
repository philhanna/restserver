package restserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Update gets an article number from the request and updates the
// corresponding article in the collection with new data. This is
// a PUT method.
func Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering Update")

	// Parse the request parameters and identify the requested ID
	id := r.URL.Query().Get("id")
	log.Printf("Updating article %q\n", id)

	// Get the body of the request (the JSON for the updated article)
	reqBody, _ := ioutil.ReadAll(r.Body)

	// Unmarshal this into a new Article struct
	var newData Article
	json.Unmarshal(reqBody, &newData)

	// Connect to the database
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update the database
	_, err = db.Exec(`
		UPDATE  articles
		SET		title = ?,
				description = ?,
				content = ?
		WHERE	id = ?		
		`, newData.Title, newData.Description, newData.Content, id)
	if err != nil {
		log.Fatal(err)
	}

	// Return the JSON representation of the new article
	json.NewEncoder(w).Encode(newData)
}
