package web_definitions

import (
	"errors"
	// "fmt"
	"log"
	"net/http"
	"strings"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// All pages are represented as structs that implement WebPageInterface and embed the PageData type
// Init should return a pointer to the page's struct
// The string parameter is the localRootFolder
// The map parameter is a map of all pages
//

type WebPageInterface interface {
	Init(string, *map[string]WebPageInterface) WebPageInterface
}

//
// All pages are represented as structs that implement WebPageInterface and embed the PageData type
// Therefore we are guaranteed that every web page has the fields defined in the PageData type below
// The Handler function implements the page's behavior
//

type PageData struct {
	Name            string                       // name
	Handler         interface{}                  // function that handles http requests. he actual implementation of page's behavior
	PageDict        *map[string]WebPageInterface // map of pointers to all other PageData structs
	UrlExtension    string                       // full url extension
	UrlSelfFolder   string                       // url extension of containing folder
	UrlRootFolder   string                       // base url extension. should be "/", "/index", or "/home/"
	UrlStaticFolder string                       // url extension of folder containing static files
	LocalRootFolder string                       // path of project's local (server) root directory
	LocalSelfFolder string                       // path of local folder containing PageData implementation and html file
	LocalHtmlFile   string                       // path of html file associated with this PageData
}

// Root folder on local machine (server)
const urlRootFolder = "/"

// ------------------------------------------- Public ------------------------------------------- //

// Creates new PageData struct and populates its fields
func NewWebPage(pageName, urlExtension, localRootFolder string, pageDict *map[string]WebPageInterface, handleFunc interface{}) *PageData {
	p := PageData{}
	p.Name = pageName
	p.UrlRootFolder = urlRootFolder
	p = *parseUrlExtension(&p, urlExtension)
	p = *parseLocalFolder(&p, localRootFolder)
	p.PageDict = pageDict

	// Ensure that handleFunc is one of the two legal types
	_, okHandler := handleFunc.(http.Handler)
	_, okHandlerFunc := handleFunc.(func(http.ResponseWriter, *http.Request))
	if !(okHandler || okHandlerFunc) {
		err := errors.New("Invalid type for handleFunc.")
		panic(err)
	}
	p.Handler = handleFunc

	return &p
}

// Wrapper for HandlerFuncs. Adds logging but will otherwise produce functionality indentical to HandlerFunc
func Logging(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// ------------------------------------------- Private ------------------------------------------- //

// All url extensions will end with "/"
func parseUrlExtension(p *PageData, urlExtension string) *PageData {

	p.UrlExtension = p.UrlRootFolder + urlExtension
	if len(urlExtension) <= 1 {
		p.UrlSelfFolder = urlExtension
	} else {
		lastSeparatorIndex := strings.LastIndex(p.UrlExtension[:len(p.UrlExtension)-1], "/")
		p.UrlSelfFolder = p.UrlExtension[:lastSeparatorIndex+1]
	}
	p.UrlSelfFolder = p.UrlExtension
	p.UrlStaticFolder = p.UrlRootFolder + "static/"

	return p
}

// All local folders will end with "\\". Do not use "/"
func parseLocalFolder(p *PageData, rootFolder string) *PageData {

	if rootFolder[len(rootFolder)-1:] != "\\" {
		rootFolder += "\\"
	}
	p.LocalRootFolder = rootFolder
	p.LocalSelfFolder = rootFolder + "template_" + p.Name + "\\"
	p.LocalHtmlFile = p.LocalSelfFolder + p.Name + ".html"

	return p
}
