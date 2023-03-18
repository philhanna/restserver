package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetAll returns the entire articles collection
func GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetAll")

	// Connect to the database
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Select all articles
	rs, err := db.Query(`

	SELECT		id,
				title,
				description,
				content
	FROM		articles
	ORDER BY 	1, 2

	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rs.Close()

	// Load the articles into an array slice
	articles := make([]Article, 0)
	for rs.Next() {
		article := Article{}
		rs.Scan(&article.Id, &article.Title, &article.Description, &article.Content)
		articles = append(articles, article)
	}

	// Return the array slice
	json.NewEncoder(w).Encode(articles)
}
