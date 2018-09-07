package main

import (
	. "Web/main_definitions"
	"Web/page_404"
	"Web/page_captcha"
	"Web/page_contact"
	"Web/page_index"
	"Web/page_resume"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// ------------------------------------------- Definitions ------------------------------------------- //

// Location of project on local machine (server)
var localRootFolder string

// Url root folder. Should be "/", "/home/", or "/main/"
const urlRootFolder = "/"

// Pointer to map of page structs. Structs must implement WebPageInterface
var pages = &map[string]WebPageInterface{
	"index":   &page_index.IndexWebPage{},
	"resume":  &page_resume.ResumeWebPage{},
	"captcha": &page_captcha.CaptchaWebPage{},
	"contact": &page_contact.ContactWebPage{},
	"404":     &page_404.FourZeroFourWebPage{},
}

// ------------------------------------------- Main ------------------------------------------- //

func init() {

	// Set up handler for serving static files
	staticFileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", staticFileServer)

	// Get project directory on local machine (server)
	localRootFolder, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	fmt.Println("LOCAL ROOT IS: " + localRootFolder)

	// Initialize (mostly) empty PageData struct
	baseData := PageData{
		LocalRootFolder: localRootFolder,
		UrlRootFolder:   urlRootFolder,
		PageDict:        pages,
	}

	// Initializes all pages by calling the Init() function of each.
	for name, emptyPage := range *pages {
		(*pages)[name] = emptyPage.Init(baseData)
	}

	// Set up the handler for each page. This can only be done when all pages are finished initializing.
	for _, page := range *pages {
		handler := page.Data().Handler
		url := page.Data().UrlExtension
		http.HandleFunc(url, handler.(func(http.ResponseWriter, *http.Request)))
	}

	fmt.Println("Finished init in main.")
	// Set up handler
}

func main() {

	fmt.Println("Running main in main.")

	// Start http server
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	// log.Fatal(http.ListenAndServe(":3000", nil))

}

// ------------------------------------------- Private ------------------------------------------- //
