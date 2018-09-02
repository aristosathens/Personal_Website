package main

import (
	// "github.com/gorilla/mux"
	// "fmt"
	// "github.com/gorilla/websocket"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
	// "path/filepath"
	"template_contact"
	"template_formatting"
	"template_index"
	"template_resume"
	. "web_definitions"
)

// ------------------------------------------- Definitions ------------------------------------------- //

// Map of pointers to web page structs. Array defined in template.go. Struct type defined in web_definitions.go
var pages *map[string]*WebPage

// Map of names to Init functions for all pages. Used to initialize pages struct above
var pageInitFunctions = map[string]func(*map[string]*WebPage) *WebPage{
	"index":      template_index.Init,
	"resume":     template_resume.Init,
	"formatting": template_formatting.Init,
	"contact":    template_contact.Init,
}

// ------------------------------------------- Main ------------------------------------------- //

func init() {

	// Set up handler for serving static files
	staticFileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	// Generate all web page objects
	pages = GetAllPages(pageInitFunctions)

	// Set up the handler for each page. This can only be done when all pages are finished initializing.
	for _, page := range *pages {
		_, isPlainHandler := page.Handler.(http.Handler)
		if isPlainHandler {
			http.Handle("/template_"+page.Name+"/", page.Handler.(http.Handler))
		} else {
			http.HandleFunc("/"+page.UrlExtension, page.Handler.(func(http.ResponseWriter, *http.Request)))
		}
	}

}

func main() {

	// Start http server
	log.Fatal(http.ListenAndServe(":3000", nil))

	// r := mux.NewRouter()

}
