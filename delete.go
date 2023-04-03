package restserver

import (
	"log"
	"net/http"
)

// Delete gets an article number from the request and deletes the
// corresponding article in the collection. This is a DELETE method.
func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering Delete")

	// Parse the request parameters and identify the requested ID
	id := r.URL.Query().Get("id")
	log.Printf("Deleting article %q\n", id)

	// Connect to the database
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Delete the article
	_, err = db.Exec(`
	
	DELETE
	FROM	articles
	WHERE	id=?
	
	`, id)
	if err != nil {
		log.Fatal(err)
	}
}
