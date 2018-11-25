package main

import (
	"net/http"

	"github.com/sonatard/appengine-go-sample/first/api"
)

func init() {
	// not logging in Stackdriver logging
	// panic("init: panic!!")

	http.HandleFunc("/", api.IndexHandle)
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
}
