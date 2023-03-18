package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// Get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	// Unmarshal this into a new Article struct
	var article Article
	json.Unmarshal(reqBody, &article)
	log.Printf("Created new article with id=%s\n", article.Id)
	
	// Append this to our articles array.
	articles = append(articles, article)

	// Respond with the ToJSON string representation of the new article.
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
