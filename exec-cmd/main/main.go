package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sonatard/appengine-go-migrate-2nd-gen-sample/exec-cmd/api"
)

func main() {
	http.HandleFunc("/", api.IndexHandle)

	port := os.Getenv("PORT")
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
