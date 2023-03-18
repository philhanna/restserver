// Tutorial source is https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// ---------------------------------------------------------------------
// Variables
// ---------------------------------------------------------------------

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

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

func handleRequests() {
	const HOST = "localhost"
	const PORT = 10000

    http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	hostAndPort := fmt.Sprintf("%s:%d", HOST, PORT)
	log.Printf("Starting server on port %d\n", PORT)
    log.Fatal(http.ListenAndServe(hostAndPort, nil))
}

func main() {
	Articles = []Article{
        {Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        {Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}