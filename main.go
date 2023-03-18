// Tutorial source is https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	const HOST = "localhost"
	const PORT = 10000
	hostAndPort := fmt.Sprintf("%s:%d", HOST, PORT)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", ReturnAllArticles)
	myRouter.HandleFunc("/article", CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", ReturnSingleArticle)

	log.Printf("Starting server on port %d\n", PORT)
	log.Fatal(http.ListenAndServe(hostAndPort, myRouter))
}

func main() {
	log.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
