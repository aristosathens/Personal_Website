package main

import (
	// "github.com/gorilla/mux"
	// "net/http"
)

// func newRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/hello", handler).Methods("GET")

// 	// Declare the static file directory and point it to the
// 	// directory we just made
// 	staticFileDirectory := http.Dir("./data/")
// 	// Declare the handler, that routes requests to their respective filename.
// 	// The fileserver is wrapped in the `stripPrefix` method, because we want to
// 	// remove the "/assets/" prefix when looking for files.
// 	// For example, if we type "/assets/index.html" in our browser, the file server
// 	// will look for only "index.html" inside the directory declared above.
// 	// If we did not strip the prefix, the file server would look for
// 	// "./assets/assets/index.html", and yield an error
// 	staticFileHandler := http.StripPrefix("/data/", http.FileServer(staticFileDirectory))
// 	// The "PathPrefix" method acts as a matcher, and matches all routes starting
// 	// with "/assets/", instead of the absolute route itself
// 	r.PathPrefix("/data/").Handler(staticFileHandler).Methods("GET")
// 	return r
// }
