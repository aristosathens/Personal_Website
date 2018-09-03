package main

import (
	"Web/template_captcha"
	"Web/template_contact"
	"Web/template_index"
	"Web/template_resume"
	. "Web/web_definitions"
	// "fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// ------------------------------------------- Definitions ------------------------------------------- //

// Location of project on local machine (server)
var localRootFolder string

// Map of pointers to web page structs. Array defined in template.go. Struct type defined in web_definitions.go
// var pages *map[string]*WebPage

// Map of empty page structs. Structs must implement WebPageInterface. Used to initialize pages struct above
var pages = &map[string]WebPageInterface{
	"index":   &template_index.IndexWebPage{},
	"resume":  &template_resume.ResumeWebPage{},
	"captcha": &template_captcha.CaptchaWebPage{},
	"contact": &template_contact.ContactWebPage{},
}

// ------------------------------------------- Main ------------------------------------------- //

func init() {

	// Set up handler for serving static files
	staticFileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", staticFileServer)

	// Get project directory on local machine (server)
	localRootFolder, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	// Generate all web page objects
	pages = GetAllPages(localRootFolder, pages)

	// Set up the handler for each page. This can only be done when all pages are finished initializing.
	for _, page := range *pages {
		data := GetData(page, "Handler", reflect.TypeOf((*string)(nil)).Elem())
		page.CaptchaLocation = data.(string)
		data := 
		data := page.GetPageData()

		_, isPlainHandler := page.GetPageData().Handler.(http.Handler)
		if isPlainHandler {
			http.Handle("/template_"+data.Name+"/", data.Handler.(http.Handler))
		} else {
			http.HandleFunc(data.UrlExtension, data.Handler.(func(http.ResponseWriter, *http.Request)))
		}
	}

}

func main() {

	// Start http server
	log.Fatal(http.ListenAndServe(":3000", nil))

	// r := mux.NewRouter()

}

// ------------------------------------------- Private ------------------------------------------- //
