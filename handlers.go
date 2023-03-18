package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateNewArticle adds a new article to the collection, based on JSON
// data contained in the request.
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering CreateNewArticle")

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

	INSERT INTO articles VALUES(?, ?, ?, ?)

	`,
		article.Id, article.Title, article.Description, article.Content)
	if err != nil {
		log.Fatal(err)
	}

	// Respond with the ToJSON string representation of the new article.
	json.NewEncoder(w).Encode(article)
}

// DeleteArticle gets an article number from the request and deletes the
// corresponding article in the collection.
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
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

	// If the article was not found, the result set will be empty.
	if !rs.Next() {
		errmsg := fmt.Sprintf("Article %s not found", id)
		http.Error(w, errmsg, http.StatusNotFound)
	}
}

// HomePage serves the home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering HomePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

// ReturnAllArticles returns the entire articles collection
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering ReturnAllArticles")

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

// ReturnSingleArticle gets an article number from the request and
// returns the corresponding article in the collection.
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering ReturnSingleArticle")

	// Get the requested Id
	vars := mux.Vars(r)
	key := vars["id"]

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

		// Otherwise, return a 404 eror
		http.Error(w, "Article not found", http.StatusNotFound)
	}
}

// UpdateArticle gets an article number from the request and updates the
// corresponding article in the collection with new data.
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering UpdateArticle")

	// Parse the request parameters and identify the requested ID
	vars := mux.Vars(r)
	id := vars["id"]

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
