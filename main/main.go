package main

import (
	"Web/template_404"
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

// Map of page structs. Structs must implement WebPageInterface
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

	// Initializes all pages by calling the Init() function of each.
	for name, emptyPage := range *pages {
		(*pages)[name] = emptyPage.Init(localRootFolder, pages)
	}

	// Set up the handler for each page. This can only be done when all pages are finished initializing.
	for _, page := range *pages {
		handler := GetData(page, "Handler", HandlerTypeArray)
		name := GetData(page, "Name", StringTypeArray).(string)
		url := GetData(page, "UrlExtension", StringTypeArray).(string)

		_, isPlainHandler := handler.(http.Handler)
		if isPlainHandler {
			http.Handle("/template_"+name+"/", handler.(http.Handler))
		} else {
			http.HandleFunc(url, handler.(func(http.ResponseWriter, *http.Request)))
		}
	}
}

func main() {

	// Start http server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
