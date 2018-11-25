package main

import (
	"net/http"

	"github.com/sonatard/appengine-go-sample/second-with-appengine-pkg/api"
	"google.golang.org/appengine"
)

/*
If dose not exist main method.
ERROR: (gcloud.app.deploy) Error Response: [9] Cloud build xxxx status: FAILURE.
Build error details: # github.com/sonatard/appengine-go-sample/second-with-appengine-pkg/src
runtime.main_main·f: function main is undeclared in the main package
*/

func main() {
	// logging in Stackdriver logging
	// panic("init: panic!!")

	http.HandleFunc("/", api.IndexHandle)
	http.HandleFunc("/cmd", api.CmdHandle)
	http.HandleFunc("/auth", api.AuthHandle)
	http.HandleFunc("/log", api.LogHandle)
	http.HandleFunc("/appenginelog", api.AppengineLogHandle)
	http.HandleFunc("/panic", api.PanicHandle)
	http.HandleFunc("/taskqueue", api.TaskQueueHandle)
	http.HandleFunc("/memcache", api.MemcacheGetHandle)
	http.HandleFunc("/memcachedeelte", api.MemcacheDeleteHandle)
	http.HandleFunc("/datastore", api.DatastoreGetHandle)
	http.HandleFunc("/datastoredelete", api.DatastoreDeleteHandle)
	http.HandleFunc("/search", api.SearchHandle)
	http.HandleFunc("/searchdelete", api.SearchDeleteHandle)

	// If not exec appengine.Main(), you can deploy but return 500
	// appengine.Main() call http.Serve and block. you have to exec last line in main method.
	appengine.Main()
	// not reach
}
