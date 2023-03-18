// Tutorial source is https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
	"log"
	"path/filepath"

	server "github.com/philhanna/restserver"
)

func main() {
	log.Println("Rest API v2.0 - Mux Routers")
	filename, err := filepath.Abs("cmd/web/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config, err := server.NewConfiguration(filename)
	if err != nil {
		log.Fatal(err)
	}
	server.HandleRequests(config)
}
