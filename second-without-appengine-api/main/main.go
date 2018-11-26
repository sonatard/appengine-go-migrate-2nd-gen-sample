package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/profiler"
	"github.com/sonatard/appengine-go-migrate-2nd-gen-sample/second-without-appengine-api/api"
)

/*
If dose not exist main method.
ERROR: (gcloud.app.deploy) Error Response: [9] Cloud build xxxx status: FAILURE.
Build error details: # github.com/sonatard/appengine-go-migrate-2nd-gen-sample/second-with-appengine-pkg/src
runtime.main_main·f: function main is undeclared in the main package
*/

func main() {
	// logging in Stackdriver logging
	// panic("init: panic!!")

	// Stackdriver Profiler
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if err := profiler.Start(profiler.Config{
		DebugLogging: false,
		ProjectID:    projectID,
	}); err != nil {
		panic(err)
	}

	http.HandleFunc("/", api.IndexHandle)
	// http.HandleFunc("/auth", api.AuthHandle)
	http.HandleFunc("/log", api.LogHandle)
	http.HandleFunc("/appenginelog", api.AppengineLogHandle)
	http.HandleFunc("/panic", api.PanicHandle)
	http.HandleFunc("/taskqueue", api.TaskQueueHandle)
	// http.HandleFunc("/memcache", api.MemcacheGetHandle)
	// http.HandleFunc("/memcachedeelte", api.MemcacheDeleteHandle)
	http.HandleFunc("/datastore", api.DatastoreGetHandle)
	http.HandleFunc("/datastoredelete", api.DatastoreDeleteHandle)
	// http.HandleFunc("/search", api.SearchHandle)
	// http.HandleFunc("/searchdelete", api.SearchDeleteHandle)

	port := os.Getenv("PORT")
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
