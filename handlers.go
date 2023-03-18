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

	// Append this to our articles array.
	articles = append(articles, article)

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

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
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
	json.NewEncoder(w).Encode(articles)
}

// ReturnSingleArticle gets an article number from the request and
// returns the corresponding article in the collection.
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering ReturnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
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

	for index, article := range articles {
		if article.Id == id {
			articles[index] = newData
		}
	}
}
