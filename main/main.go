package main

import (
	"Web/template_404"
	"Web/template_captcha"
	"Web/template_contact"
	"Web/template_index"
	"Web/template_resume"
	. "Web/web_definitions"
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
	"index":   &template_index.IndexWebPage{},
	"resume":  &template_resume.ResumeWebPage{},
	"captcha": &template_captcha.CaptchaWebPage{},
	"contact": &template_contact.ContactWebPage{},
	"404":     &template_404.FourZeroFourWebPage{},
}

// ------------------------------------------- Main ------------------------------------------- //

func init() {

	// Set up handler for serving static files
	staticFileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", staticFileServer)

	// Get project directory on local machine (server)
	localRootFolder, _ = filepath.Abs(filepath.Dir(os.Args[0]))

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

	// Set up handler
}

func main() {

	// Start http server
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// ------------------------------------------- Private ------------------------------------------- //
