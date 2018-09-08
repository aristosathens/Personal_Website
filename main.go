package main

import (
	. "Web/main_definitions"
	"Web/page_404"
	"Web/page_captcha"
	"Web/page_contact"
	"Web/page_index"
	"Web/page_projects"
	"Web/page_resume"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// ------------------------------------------- Definitions ------------------------------------------- //

// Location of project on local machine (server)
var localRootFolder string

// Url root folder. Should be "/", "/home/", or "/main/"
const urlRootFolder = "/"

// Pointer to map of page structs. Structs must implement WebPageInterface
var pages = &map[string]WebPageInterface{
	"index":    &page_index.IndexWebPage{},
	"resume":   &page_resume.ResumeWebPage{},
	"captcha":  &page_captcha.CaptchaWebPage{},
	"contact":  &page_contact.ContactWebPage{},
	"404":      &page_404.FourZeroFourWebPage{},
	"projects": &page_projects.ProjectsWebPage{},
}

// ------------------------------------------- Main ------------------------------------------- //

func init() {

	// Set up handler for serving static files
	serveAllStatic()

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
}

func main() {

	fmt.Println("Website now running.")
	// Where are we running? (windows == my local machine)
	var port string
	if runtime.GOOS == "windows" {
		port = ":3000"
	} else {
		port = ":" + os.Getenv("PORT)")
	}

	// Start http server
	log.Fatal(http.ListenAndServe(port, nil))

}

// ------------------------------------------- Private ------------------------------------------- //

// Serves all files and subdirectories in local static/ folder
func serveAllStatic() {

	// Serve all files in local static/ folder
	staticFileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", staticFileServer)

	// Serve all subdirectories in local static/ folder
	subDir, _ := ioutil.ReadDir("static")
	for _, f := range subDir {
		if f.Mode().IsDir() {
			dirName := "/static/" + f.Name()
			staticDirServer := http.StripPrefix(dirName, http.FileServer(http.Dir(dirName[1:len(dirName)-1])))
			http.Handle(dirName, staticDirServer)
		}
	}
}
