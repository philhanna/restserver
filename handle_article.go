package restserver

import (
	"log"
	"net/http"
)

func HandleArticle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request method is %q\n", r.Method)
	switch r.Method {
	default:
		Get(w, r)
	case "PUT":
		Update(w, r)
	case "DELETE":
		Delete(w, r)
	case "POST":
		Create(w, r)
	}
}
