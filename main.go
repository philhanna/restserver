// Tutorial source is https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// ---------------------------------------------------------------------
// Variables
// ---------------------------------------------------------------------



// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    log.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnSingleArticle")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	const HOST = "localhost"
	const PORT = 10000
	hostAndPort := fmt.Sprintf("%s:%d", HOST, PORT)

	myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	
	log.Printf("Starting server on port %d\n", PORT)
    log.Fatal(http.ListenAndServe(hostAndPort, myRouter))
}

func main() {
	log.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}