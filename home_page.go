package main

import (
	"fmt"
	"log"
	"net/http"
)

// HomePage serves the home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering HomePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
