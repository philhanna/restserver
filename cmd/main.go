// Tutorial source is https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
	"log"
	server "github.com/philhanna/restserver"
)

func main() {
	log.Println("Rest API v2.0 - Mux Routers")
	server.HandleRequests()
}
